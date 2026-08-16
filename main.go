package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"html/template"
) 

var homeTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/home.html"))
var projectsTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/projects.html"))
var projectDetailTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/project_detail.html"))
var contactTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/contact.html"))
var resumeTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/resume.html"))
var blogTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/blog.html"))
var askTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/ask.html"))


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

var contact = Contact{
	Github: "http://github.com/leouduh",
	Email: "mailto:chigozieuduh.cu@gmail.com",
	Discord: "https://discord.com/users/852747108212539422",

}

var github = Github{
	GithubUrl: "https://github.com/leouduh",
}

var experience = []Experience{
	{
		Role: "Machine Learing/MLOps Egineer",
		Company: "Jaguar Land Rover",
		Period: "May 2023 - Present",
		Highlights: []string{
			"Deployed a NLP service in production serving thousands of customers",
		},
	},
	{
		Role: "Software Engineer",
		Company: "Hauwei Technologies Co., Ltd.",
		Period: "February 2021 - December 2021",
		Highlights: []string{
			"Built web applications for Internet Service Providers like MTN and Airtel in SubSaharan Africa",
		},
	},
}

var aboutLeo string ="I am MLOps Engineer that operates at the intersection between data science, machine learning and software engineering, experienced at deploying ml models into production on cloud infrastructures like AWS"


var skills = []string{
	"Python",
	"CI/CD Gitlab",
	"Docker",
	"AWS CDK",
	"AWS",
	"Terraform",
	"Model Deployment",
	"Model Monitoring and Observability",
	"Golang",
	"C programming language",
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	data := Page{Title: "Home", Skills: skills, AboutLeo: aboutLeo,}
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
		Github: github,
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
	data := Page{
		Title: "resume",
		Experience: experience,
		Contact: contact,
	}
	err := resumeTmpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	data := Page{
		Title: "contact-info",
		Github: github,
		Contact: contact,
	}
	err := contactTmpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func blogHangler(w http.ResponseWriter, r *http.Request){
	data := Page{
		Title: "blog",
	}
	err := blogTmpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func askHandler(w http.ResponseWriter, r *http.Request){
	data := Page{
		Title: "ask-ai",
	}
	err := askTmpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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
	mux.HandleFunc("/blog", blogHangler)
	mux.HandleFunc("/ask", askHandler)

	addr := ":" + port
	log.Printf("Server starting on address %s", addr)

	err := http.ListenAndServe(addr, stripTrailingSlash(mux))
	if err != nil{
		log.Fatalf("server failed: %v", err)
	}

}
