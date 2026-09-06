package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
	"os"
	"strings"
)

var db *sql.DB

func initDB() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./visits.db"
	}
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("An error was encountered trying to open DB: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS visits (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip_hash TEXT NOT NULL,
		path TEXT NOT NULL,
		ts DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("Error in executing schema: %v", err)
	}

	log.Printf("visits database is ready at %s", dbPath)

}

// This function just used for a one way hash to hash ip address and return string. each visitor will have a unique hash
func hashIP(ip string) string {
	digest := sha256.Sum256([]byte(ip))
	return hex.EncodeToString(digest[:])

}

func trackVisits(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return

		}

		hashed := hashIP(clientIP(r))
		if _, err := db.Exec(
			"INSERT INTO visits (ip_hash, path) VALUES (?, ?)", hashed, r.URL.Path,
		); err != nil {
			log.Printf("failed to record visit: %v", err)
		}

		next.ServeHTTP(w, r)
	})

}
