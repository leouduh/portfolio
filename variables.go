package main

var projects = []Project{
	{
		Slug:        "porfolio-website",
		Title:       "leo's portfolio website",
		Summary:     "Personal portfolio webiste where I showcase myself to the rest of the world",
		Description: "This is something I have been putting off and procrastinating on for so many years finally " +
		"getting this done here. I show case in this website how I use golang, css, html to get something up and " +
		"running, I will be adding blogs, and my thoughts here. I hope this will not be an abandoned project and " + 
		"I can keep working on it. This website features a little chatbot where people can ask questions about me if" +
		"they don't want to reach me",
		Stack:       []string{"Golang", "html", "CSS"},
		Featured:    true,
	},
	// {Slug: "", Title: "", Summary: "", Featured: true},

}

var contact = Contact{
	Github:  "http://github.com/leouduh",
	Email:   "mailto:chigozieuduh.cu@gmail.com",
	Discord: "https://discord.com/users/852747108212539422",
}

var github = Github{
	GithubUrl: "https://github.com/leouduh",
}

var experience = []Experience{
	{
		Role:    "Machine Learing/MLOps Egineer",
		Company: "Jaguar Land Rover",
		Period:  "May 2023 - Present",
		Highlights: []string{
			"Deployed a NLP service in production serving thousands of customers",
		},
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

var p1 string = "Funny thing my first name is actually Chigozie but I go by Leo professionally," +
	"everyone pretty much calls me leo. I build the IaC and pipelines that help data scientists get their " +
	"beautiful not so boring projects out of notebooks and into cloud environments useful for end users and " +
	"internal stakeholders. At JLR I have shipped a bunch of data and ml pipelines and NLP recommendation service " +
	"in production serving thousands of customers, before that I built web platforms for telecom providers across " +
	"Sub-Saharan African back in my home country-Nigeria."
var p2 string = "Outside of work, I enjoy playing sports football (it is football and not soccer btw) and getting " +
	"active either going to the gym or on some hike in the landscapes of ireland. Going to be using this space to " +
	"Document finding and bloging about randome stuff I find interesting in tech and outside of tecch"
var p3 string = "Working on something intersing or just wnat to say hi? reach out to me or ask that little ai chatbot"
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
}
