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
	
	// mux is a ServeMux - a lookup table mapping URL patterns to handlers. For every request it finds the
	// most specific matching pattern and calls that pattern's handler.
	mux := http.NewServeMux()

	// fileServer knows how to read files off disk and write them into an HTTP response. http.Dir("./static/")
	// tells it: "your root folder is ./static/ on this computer's disk." This path is never seen by the
	// browser - it has nothing to do with URLs, it's purely where I look on THIS machine.
	fileServer := http.FileServer(http.Dir("./static/"))

	// Registered with a trailing slash ("/static/"), so this is a SUBTREE match - it catches this exact
	// path AND anything nested under it, at any depth (/static/css/output.css, /static/img/icons/x.svg,
	// etc). Compare to exact-match routes like "/resume" below (no trailing slash) - those match ONLY
	// that literal path, nothing nested under them at all.
	//
	// StripPrefix removes the "/static/" part of the URL before fileServer ever sees it. Without this,
	// fileServer would glue the FULL incoming path onto its root and look for
	// ./static/static/css/output.css (folder name doubled, doesn't exist). With it, fileServer only ever
	// sees "css/output.css" and correctly finds ./static/css/output.css.
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("GET /projects/{slug}", projectDetailHandler)
	// No trailing slash = EXACT match only. Matches "/resume" and nothing else - not "/resume/", not
	// "/resume/anything". Deliberate: there's no legitimate sub-page under /resume, so exact-match is the
	// correct, intentional choice here (unlike /static/ above, which genuinely needs arbitrarily nested
	// paths to work). Same reasoning applies to /contact, /blog, and /ask below.
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
