package work

import (
	"net/http"

	"mu.dev"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	html := mu.Template("Work", "Do good work", "", `
	<h1 style="padding-top: 50px;">Find Work</h1>

	No jobs yet

	<a href="mailto:contact@mu.xyz">Post a job</a>
	`)
}

func Register() {}
