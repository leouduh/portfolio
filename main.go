package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"html/template"
) 

var homeTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))

func homeHandler(w http.ResponseWriter, r *http.Request){
	type Page struct{
		Title string
	}
	data := Page{Title: "Home"}
	err := homeTmpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "temporary projects landing page")

}

func resumeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "temporary resume landing page")

}

func contactHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "temporary contact landing page")
}

// write middle ware function to intercept trailing slashes
// you pass the next function, which is the function that will run afte rthe middle ware function
func stripTrailingSlash(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// if block to handle if it has a trailing "/"
		if len(r.URL.Path) > 1 && strings.HasSuffix(r.URL.Path, "/"){
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("/resume", resumeHandler)
	mux.HandleFunc("/contact", contactHandler)

	addr := ":" + port
	log.Printf("Server starting on address %s", addr)

	err := http.ListenAndServe(addr, stripTrailingSlash(mux))
	if err != nil{
		log.Fatalf("server failed: %v", err)
	}

}
