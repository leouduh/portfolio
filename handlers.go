package main

import (
	"html/template"
	"net/http"
)

var homeTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
var projectsTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/projects.html"))
var projectDetailTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/project_detail.html"))
var resumeTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/resume.html"))
var blogTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/blog.html"))
var askTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/ask.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title:    "Home",
		Skills:   skills,
		AboutLeo: aboutLeo,
		Contact:  contact,
	}
	err := homeTmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	var featured []Project
	for _, p := range projects {
		if p.Featured {
			featured = append(featured, p)
		}
	}
	data := Page{
		Title:    "Projects",
		Projects: featured,
		Github:   github,
		Contact:  contact,
	}
	err := projectsTmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func projectDetailHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	// iterate through all projects defined outside main
	for _, p := range projects {
		// if the project is the specific one that the user requested for process accordingly
		if p.Slug == slug {
			data := Page{
				Title:   p.Title,
				Project: p,
			}
			err := projectDetailTmpl.ExecuteTemplate(w, "base", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.NotFound(w, r)
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title:      "resume",
		Experience: experience,
		Contact:    contact,
	}
	err := resumeTmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func blogHangler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title:   "blog",
		Contact: contact,
	}
	err := blogTmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func askHandler(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title: "ask-ai",
	}
	err := askTmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
