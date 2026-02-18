package web

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"

	"go-passman/internal/models"
	"go-passman/internal/storage"
)

//go:embed templates/*
var templatesFS embed.FS

var (
	tmpl       *template.Template
	vault      *models.Vault
	vaultPwd   *string
	vaultPwdStored string // keep password on heap so vaultPwd stays valid
	vaultMu    sync.RWMutex
)

func init() {
	tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"urlquery": url.QueryEscape,
		"add":      func(a, b int) int { return a + b },
		"sub":      func(a, b int) int { return a - b },
		"iterate":  func(start, end int) (out []int) {
			for i := start; i <= end; i++ {
				out = append(out, i)
			}
			return out
		},
		"inactivityMinutes": func() int {
			if s := os.Getenv("INACTIVITY_MINUTES"); s != "" {
				if n, err := strconv.Atoi(s); err == nil && n > 0 {
					return n
				}
			}
			return 5
		},
	}).ParseFS(templatesFS, "templates/*.html"))
}

// loadVault loads vault into memory (with optional password for encrypted vault).
func loadVault(w http.ResponseWriter, r *http.Request) (*models.Vault, *string, bool) {
	vaultMu.RLock()
	v, p := vault, vaultPwd
	vaultMu.RUnlock()
	if v != nil {
		return v, p, true
	}
	enc, err := storage.IsVaultEncrypted()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, nil, false
	}
	if enc {
		http.Redirect(w, r, "/unlock", http.StatusFound)
		return nil, nil, false
	}
	v, p, err = storage.LoadVaultWithPassword(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, nil, false
	}
	vaultMu.Lock()
	vault, vaultPwd = v, p
	vaultMu.Unlock()
	return v, p, true
}

// Run starts the web server on 127.0.0.1:8080 (use WEB_PORT env to override port).
func Run() {
	addr := "127.0.0.1:8080"
	if p := os.Getenv("WEB_PORT"); p != "" {
		addr = "127.0.0.1:" + p
	}
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/api/copy", copyHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/unlock", unlockHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/show", showHandler)
	log.Printf("go-passman web UI: http://%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
