package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

// write middle ware function to intercept trailing slashes
// you pass the next function, which is the function that will run afte rthe middle ware function
func stripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if block to handle if it has a trailing "/"
		if len(r.URL.Path) > 1 && strings.HasSuffix(r.URL.Path, "/") {
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("GET /projects/{slug}", projectDetailHandler)
	mux.HandleFunc("/resume", resumeHandler)
	mux.HandleFunc("/contact", contactHandler)
	mux.HandleFunc("/blog", blogHangler)
	mux.HandleFunc("/ask", askHandler)

	addr := ":" + port
	log.Printf("Server starting on address %s", addr)

	err := http.ListenAndServe(addr, stripTrailingSlash(mux))
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
