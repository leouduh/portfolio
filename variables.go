package main

var projects = []Project{
	{
		Slug:    "porfolio-website",
		Title:   "My portfolio website",
		Summary: "Personal portfolio webiste where I showcase myself to the rest of the world",
		Description: "This is something I have been putting off and procrastinating on for so many years finally " +
			"getting this done here. I show case in this website how I use golang, css, html to get something up and " +
			"running, I will be adding blogs, and my thoughts here. I hope this will not be an abandoned project and " +
			"I can keep working on it. This website features a little chatbot where people can ask questions about me if" +
			"they don't want to reach me",
		Stack:    []string{"Golang", "html", "CSS", "JavaScript", "Terraform"},
		Featured: true,
	},
	{
		Slug: "ityb",
		Title: "I Type you Build",
		Summary: "",
		Description: "",
		Stack: []string{"Python"},
		Featured: true,
	},

}

var contact = Contact{
	Github:    "http://github.com/leouduh",
	Email:     "mailto:chigozieuduh.cu@gmail.com",
	Discord:   "https://discord.com/users/852747108212539422",
	Instagram: "https://www.instagram.com/directurban",
	LinkedIn:  "https://www.linkedin.com/in/chigozie-leo-uduh-42883114a/",
}

var github = Github{
	GithubUrl: "https://github.com/leouduh",
}

var jlr_experiences = []string{
	"Deployed JLR's first ML inference model on the public internet, powering NLP-based vehicle recommendations on rangerover.com",
	"Owned infrastructure, delivery, and operations end-to-end for the chapter's first production AWS deployment",
	"Built the chapter's monitoring/observability baseline and release process — immutable artefacts, smoke tests, canary rollbacks",
	"Main contributor to the chapter's internal AWS CDK library, letting data scientists deploy infra via config instead of code",
	"Added Iceberg table management, canary deployments, and observability-by-default to the CDK library",
	"Delivered all of the above inside strict GDPR compliance boundaries",
}

var h_experiences = []string{
	"",
}

var unecon_experiences = []string{
	"",
}

var experience = []Experience{
	{
		Role:       "Machine Learning/MLOps Engineer",
		Company:    "Jaguar Land Rover",
		Period:     "May 2023 - Present",
		Highlights: jlr_experiences,
	},
	{
		Role:    "Software Engineer",
		Company: "Hauwei Technologies Co., Ltd.",
		Period:  "February 2021 - December 2021",
		Highlights: []string{
			"Built web applications for Internet Service Providers like MTN and Airtel in SubSaharan Africa",
		},
	},
}

var p1 string = "Fun fact: my actual first name is Chigozie — Leo is my middle name, but it's the " +
	"one that stuck, so that's what everyone calls me. I build the IaC and pipelines that help data " +
	"scientists turn their beautiful, not-so-boring notebook projects into things that actually run in " +
	"the cloud — specifically AWS. I basically breathe life into their POCs. At JLR, I've shipped data " +
	"and ML pipelines specifically around anomaly detection, saving the company real money in warranty " +
	"claims, along with an NLP recommendation service currently serving thousands of customers in " +
	"production. Before that, I built web platforms for telecom providers across Sub-Saharan Africa, " +
	"back in my home country — Nigeria."
var p2 string = "Outside of work, I play basketball and football (football, not soccer, thank you) " +
	"and stay active with the gym or a good hike. I've recently picked up bouldering and I'm hooked. " +
	"I'll be using this space to document findings and blog about random stuff I find interesting, in " +
	"tech and out of it."
var p3 string = "Working on something interesting, or just want to say hi? Reach out — or ask that " +
	"little AI chatbot about me instead."
var aboutLeo = []string{
	p1,
	p2,
	p3,
}
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
	"Kubernetes",
	"Snowflake",
}
