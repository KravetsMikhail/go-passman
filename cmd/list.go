package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"go-passman/internal/models"
	"go-passman/internal/storage"
	"go-passman/internal/utils"
	"golang.org/x/term"
)

const (
	// Compact max lengths so one line fits in ~70 chars and doesn't require wide terminal
	maxServiceLen = 14
	maxLoginLen   = 14
	maxHostLen    = 20
	maxCommentLen = 22
	sep            = " · "

	listPageSize = 20 // entries per page when listing (paginated if total > listPageSize)
)

// NewListCommand creates the list command
func NewListCommand() *cobra.Command {
	var table bool
	var filter string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all services or entries in the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleList(table, filter)
		},
	}

	cmd.Flags().BoolVarP(&table, "table", "t", false, "Show as table (wide; use when terminal is wide enough)")
	cmd.Flags().StringVarP(&filter, "filter", "f", "", "Show only entries whose name contains this substring (case-insensitive)")

	return cmd
}

func truncate(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

// getSortedServices returns service names in stable sorted order (used for list numbering and copy/remove by number).
func getSortedServices(entries map[string]models.PasswordEntry) []string {
	services := make([]string, 0, len(entries))
	for s := range entries {
		services = append(services, s)
	}
	sort.Strings(services)
	return services
}

// resolveServiceOrNumber resolves argument to a service name: "2" -> Nth entry (1-based), "github" -> "github".
func resolveServiceOrNumber(entries map[string]models.PasswordEntry, input string) (string, error) {
	services := getSortedServices(entries)
	if n, err := strconv.Atoi(strings.TrimSpace(input)); err == nil {
		if n < 1 || n > len(services) {
			return "", fmt.Errorf("number %d out of range (1-%d)", n, len(services))
		}
		return services[n-1], nil
	}
	name := input
	if _, exists := entries[name]; exists {
		return name, nil
	}
	return "", fmt.Errorf("service '%s' not found", name)
}

func handleList(useTable bool, filter string) error {
	vault, _, err := storage.LoadVault()
	if err != nil {
		return err
	}

	if len(vault.Entries) == 0 {
		fmt.Println("📭 No passwords saved yet.")
		return nil
	}

	allServices := make([]string, 0, len(vault.Entries))
	for s := range vault.Entries {
		allServices = append(allServices, s)
	}
	sort.Strings(allServices)

	services := allServices
	var numbers []int // 1-based display numbers; when set, use for "copy N" consistency (e.g. with filter)
	if filter != "" {
		sub := strings.TrimSpace(strings.ToLower(filter))
		filtered := make([]string, 0)
		numFiltered := make([]int, 0)
		for i, s := range allServices {
			if strings.Contains(strings.ToLower(s), sub) {
				filtered = append(filtered, s)
				numFiltered = append(numFiltered, i+1)
			}
		}
		services = filtered
		numbers = numFiltered
		if len(services) == 0 {
			fmt.Printf("📭 No entries match filter %q.\n", filter)
			return nil
		}
	}

	total := len(services)
	totalInVault := len(allServices)
	stdoutTTY := term.IsTerminal(int(os.Stdout.Fd()))
	if total <= listPageSize || !stdoutTTY {
		// No pagination: small list or output redirected (e.g. list > file)
		if useTable {
			printListTableWithNumbers(services, numbers, vault.Entries, total, totalInVault)
		} else {
			printListCompactWithNumbers(services, numbers, vault.Entries, total, totalInVault)
		}
		return nil
	}

	// Paginated output (interactive terminal)
	for start := 0; start < total; start += listPageSize {
		end := start + listPageSize
		if end > total {
			end = total
		}
		page := services[start:end]
		var pageNumbers []int
		if numbers != nil {
			pageNumbers = numbers[start:end]
		} else {
			pageNumbers = nil
		}
		pageNum := start/listPageSize + 1
		totalPages := (total + listPageSize - 1) / listPageSize
		fmt.Printf("🔐 Saved entries (page %d of %d). Найдено %d из %d\n", pageNum, totalPages, total, totalInVault)
		fmt.Println()
		if useTable {
			printListTablePageWithNumbers(page, pageNumbers, vault.Entries)
		} else {
			printListCompactPageWithNumbers(page, pageNumbers, vault.Entries)
		}
		if end < total {
			line, _ := utils.ReadInput("Press Enter for next page (q to quit): ")
			if strings.TrimSpace(strings.ToLower(line)) == "q" {
				break
			}
		}
	}
	return nil
}

// printListCompact prints one short line per entry with number. Fits narrow terminals; lines may wrap but don't slide.
func printListCompact(services []string, entries map[string]models.PasswordEntry) {
	n := len(services)
	printListCompactWithNumbers(services, nil, entries, n, n)
}

// printListCompactWithNumbers prints one short line per entry; numbers are 1-based (nil = 1,2,3...). shown and total are for "найдено X из Y".
func printListCompactWithNumbers(services []string, numbers []int, entries map[string]models.PasswordEntry, shown, total int) {
	fmt.Printf("🔐 Saved entries (copy N; run 'remove' to delete). Найдено %d из %d\n", shown, total)
	fmt.Println()
	printListCompactPageWithNumbers(services, numbers, entries)
}

// printListCompactPageWithNumbers prints one page; numbers[i] is the display number (nil = 1,2,...).
func printListCompactPageWithNumbers(services []string, numbers []int, entries map[string]models.PasswordEntry) {
	for i, service := range services {
		num := i + 1
		if numbers != nil && i < len(numbers) {
			num = numbers[i]
		}
		entry := entries[service]
		s := truncate(service, maxServiceLen)
		l := truncate(entry.Login, maxLoginLen)
		h := truncate(entry.Host, maxHostLen)
		c := truncate(entry.Comment, maxCommentLen)
		if l == "" {
			l = "-"
		}
		if h == "" {
			h = "-"
		}
		if c == "" {
			c = "-"
		}
		line := fmt.Sprintf("%d.", num) + "  " + s + sep + l + sep + h + sep + c
		fmt.Println("  " + line)
	}
	fmt.Println()
}

// printListTable prints aligned columns with # column. Use only when terminal is wide (e.g. go-passman list -t).
func printListTable(services []string, entries map[string]models.PasswordEntry) {
	n := len(services)
	printListTableWithNumbers(services, nil, entries, n, n)
}

// printListTableWithNumbers prints table; numbers are 1-based (nil = 1,2,3...). shown and total are for "найдено X из Y".
func printListTableWithNumbers(services []string, numbers []int, entries map[string]models.PasswordEntry, shown, total int) {
	fmt.Printf("🔐 Saved entries (copy N; run 'remove' to delete). Найдено %d из %d\n", shown, total)
	fmt.Println()
	printListTablePageWithNumbers(services, numbers, entries)
}

// printListTablePageWithNumbers prints one page of table; numbers[i] is the display number (nil = 1,2,...).
func printListTablePageWithNumbers(services []string, numbers []int, entries map[string]models.PasswordEntry) {
	widths := [5]int{4, maxServiceLen, maxLoginLen, maxHostLen, maxCommentLen}
	headers := [5]string{"#", "Service", "Login", "Host", "Comment"}

	fmt.Println("  " + tableRow5(headers, widths))
	fmt.Println("  " + strings.Repeat("-", widths[0]+widths[1]+widths[2]+widths[3]+widths[4]+8))

	for i, service := range services {
		num := i + 1
		if numbers != nil && i < len(numbers) {
			num = numbers[i]
		}
		entry := entries[service]
		row := [5]string{
			fmt.Sprintf("%d", num),
			truncate(service, widths[1]),
			truncate(entry.Login, widths[2]),
			truncate(entry.Host, widths[3]),
			truncate(entry.Comment, widths[4]),
		}
		fmt.Println("  " + tableRow5(row, widths))
	}
	fmt.Println()
}

func tableRow5(fields [5]string, widths [5]int) string {
	var b strings.Builder
	for i := 0; i < 5; i++ {
		if i > 0 {
			b.WriteString("  ")
		}
		s := fields[i]
		if len(s) > widths[i] {
			s = s[:widths[i]-3] + "..."
		}
		b.WriteString(s)
		b.WriteString(strings.Repeat(" ", widths[i]-len(s)))
	}
	return b.String()
}
