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
var projectsTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/projects.html"))
var projectDetailTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/project_detail.html"))

type Project struct {
	Slug string
	Title string
	Summary string
	Description string
	Stack []string
	RepoURL string
	Featured bool
}

type Page struct{
	Title string
	Projects []Project
	Project Project
	Github string
}

var projects = []Project{
	{
		Slug: "porfolio-website",
		Title: "leo's portfolio website",
		Summary: "Personal portfolio webiste where I showcase myself to the rest of the world", 
		Description: "This is something I have been putting off and procrastinating on for so many years finally getting this done here. I show case in this website how I use golang, css, html to get something up and running, I will be adding blogs, and my thoughts here. I hope this will not be an abandoned project and I can keep working on it. This website features a little chatbot where people can ask questions about me if they don't want to reach me",
		Stack: []string{"Golang", "html", "CSS"},
		Featured: true,
	},
	// {Slug: "", Title: "", Summary: "", Featured: true},

}
func homeHandler(w http.ResponseWriter, r *http.Request){
	data := Page{Title: "Home"}
	err := homeTmpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request){
	var featured []Project
	for _, p := range projects{
		if p.Featured {
			featured = append(featured, p)
		}
	}
	data := Page{
		Title: "Projects",
		Projects: featured, 
		Github: "https://github.com/leouduh",
	}
	err := projectsTmpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func projectDetailHandler(w http.ResponseWriter, r *http.Request){
	slug := r.PathValue("slug")
	// iterate through all projects defined outside main
	for _, p := range projects{
		// if the project is the specific one that the user requested for process accordingly
		if p.Slug == slug{
			data := Page{
				Title: p.Title,
				Project: p,
			}
			err := projectDetailTmpl.ExecuteTemplate(w, "base", data)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.NotFound(w, r)
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
	mux.HandleFunc("GET /projects/{slug}", projectDetailHandler)
	mux.HandleFunc("/resume", resumeHandler)
	mux.HandleFunc("/contact", contactHandler)

	addr := ":" + port
	log.Printf("Server starting on address %s", addr)

	err := http.ListenAndServe(addr, stripTrailingSlash(mux))
	if err != nil{
		log.Fatalf("server failed: %v", err)
	}

}
