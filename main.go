package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
)

// Уязвимость 1: Утечка секрета (поймает Gitleaks)
const AWSAccessKeyId = "AKIAIOSFODNN7EXAMPLE"

func main() {
	// Уязвимость 2: Command Injection (поймает Semgrep)
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		target := r.URL.Query().Get("target")
		cmd := exec.Command("sh", "-c", "ping -c 1 "+target)
		out, _ := cmd.CombinedOutput()
		fmt.Fprintf(w, "%s", out)
	})

	// Уязвимость 3: SQL Injection (поймает Semgrep)
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		db, _ := sql.Open("postgres", "postgres://user:password@localhost/db")
		query := fmt.Sprintf("SELECT username FROM users WHERE id = '%s'", id)
		db.Query(query)
		fmt.Fprintf(w, "User requested")
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
