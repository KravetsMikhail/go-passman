package web

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"go-passman/internal/models"
	"go-passman/internal/storage"
)

const listPerPage = 20

type listData struct {
	Entries      []listEntry
	Total        int   // total in vault
	TotalFiltered int   // after search
	Query        string
	Page         int
	TotalPages   int
	PerPage      int
}

type listEntry struct {
	Num     int
	Name    string
	Login   string
	Host    string
	Comment string
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	v, _, ok := loadVault(w, r)
	if !ok {
		return
	}
	names := make([]string, 0, len(v.Entries))
	for n := range v.Entries {
		names = append(names, n)
	}
	sort.Strings(names)
	all := make([]listEntry, 0, len(names))
	for i, name := range names {
		e := v.Entries[name]
		all = append(all, listEntry{
			Num:     i + 1,
			Name:    name,
			Login:   e.Login,
			Host:    e.Host,
			Comment: e.Comment,
		})
	}

	query := strings.TrimSpace(r.URL.Query().Get("q"))
	filtered := all
	if query != "" {
		q := strings.ToLower(query)
		filtered = make([]listEntry, 0)
		for i, e := range all {
			if strings.Contains(strings.ToLower(e.Name), q) ||
				strings.Contains(strings.ToLower(e.Login), q) ||
				strings.Contains(strings.ToLower(e.Host), q) ||
				strings.Contains(strings.ToLower(e.Comment), q) {
				filtered = append(filtered, listEntry{
					Num: i + 1, Name: e.Name, Login: e.Login, Host: e.Host, Comment: e.Comment,
				})
			}
		}
		// renumber filtered list 1..n
		for i := range filtered {
			filtered[i].Num = i + 1
		}
	}

	totalFiltered := len(filtered)
	totalPages := 1
	if totalFiltered > listPerPage {
		totalPages = (totalFiltered + listPerPage - 1) / listPerPage
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	start := (page - 1) * listPerPage
	end := start + listPerPage
	if end > totalFiltered {
		end = totalFiltered
	}
	pageEntries := filtered[start:end]
	// display numbers on page (global index in filtered list)
	for i := range pageEntries {
		pageEntries[i].Num = start + i + 1
	}

	tmpl.ExecuteTemplate(w, "list.html", listData{
		Entries:       pageEntries,
		Total:         len(all),
		TotalFiltered: totalFiltered,
		Query:         query,
		Page:         page,
		TotalPages:   totalPages,
		PerPage:      listPerPage,
	})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	vaultMu.Lock()
	vault = nil
	vaultPwd = nil
	vaultPwdStored = ""
	vaultMu.Unlock()
	http.Redirect(w, r, "/unlock", http.StatusFound)
}

func unlockHandler(w http.ResponseWriter, r *http.Request) {
	vaultMu.RLock()
	hasVault := vault != nil
	vaultMu.RUnlock()
	if hasVault {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		pwd := r.FormValue("password")
		if pwd == "" {
			tmpl.ExecuteTemplate(w, "unlock.html", "Password required")
			return
		}
		v, _, err := storage.LoadVaultWithPassword(&pwd)
		if err != nil {
			tmpl.ExecuteTemplate(w, "unlock.html", "Wrong password or error: "+err.Error())
			return
		}
		vaultMu.Lock()
		vault = v
		vaultPwdStored = pwd
		vaultPwd = &vaultPwdStored
		vaultMu.Unlock()
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tmpl.ExecuteTemplate(w, "unlock.html", nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	v, pwd, ok := loadVault(w, r)
	if !ok {
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := strings.TrimSpace(r.FormValue("name"))
		if name == "" {
			tmpl.ExecuteTemplate(w, "add.html", "Service name is required")
			return
		}
		if _, exists := v.Entries[name]; exists {
			tmpl.ExecuteTemplate(w, "add.html", "Service already exists")
			return
		}
		v.Entries[name] = models.PasswordEntry{
			Login:     strings.TrimSpace(r.FormValue("login")),
			Host:      strings.TrimSpace(r.FormValue("host")),
			Comment:   strings.TrimSpace(r.FormValue("comment")),
			Password:  r.FormValue("password"),
			Encrypted: false,
		}
		if err := storage.SaveVault(v, pwd); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tmpl.ExecuteTemplate(w, "add.html", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	v, pwd, ok := loadVault(w, r)
	if !ok {
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	entry, exists := v.Entries[name]
	if !exists {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		newName := strings.TrimSpace(r.FormValue("name"))
		if newName == "" {
			newName = name
		}
		entry.Login = strings.TrimSpace(r.FormValue("login"))
		entry.Host = strings.TrimSpace(r.FormValue("host"))
		entry.Comment = strings.TrimSpace(r.FormValue("comment"))
		if p := r.FormValue("password"); p != "" {
			entry.Password = p
		}
		if newName != name {
			delete(v.Entries, name)
		}
		v.Entries[newName] = entry
		if err := storage.SaveVault(v, pwd); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	data := map[string]interface{}{
		"Name":    name,
		"Login":   entry.Login,
		"Host":    entry.Host,
		"Comment": entry.Comment,
		"Error":   nil,
	}
	tmpl.ExecuteTemplate(w, "edit.html", data)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	v, pwd, ok := loadVault(w, r)
	if !ok {
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if r.Method == http.MethodPost {
		delete(v.Entries, name)
		if err := storage.SaveVault(v, pwd); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	tmpl.ExecuteTemplate(w, "delete.html", name)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	v, _, ok := loadVault(w, r)
	if !ok {
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	entry, exists := v.Entries[name]
	if !exists {
		http.NotFound(w, r)
		return
	}
	tmpl.ExecuteTemplate(w, "show.html", map[string]string{"Name": name, "Password": entry.Password})
}
