package main

type Project struct {
	Slug        string
	Title       string
	Summary     string
	Description string
	Stack       []string
	RepoURL     string
	Featured    bool
}

type Contact struct {
	Email   string
	Discord string
	Github  string
}

type Github struct {
	GithubUrl string
}

type Experience struct {
	Role       string
	Company    string
	Period     string
	Highlights []string
}
type Page struct {
	Title      string
	Projects   []Project
	Project    Project
	Github     Github
	Contact    Contact
	Experience []Experience
	Skills     []string
	AboutLeo   []string
}
