package work

import (
	"net/http"

	"mu.dev"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	html := mu.Template("Work", "Do work, get paid", "", `
	<h1 style="padding-top: 100px">Work</h1>
	<p>No jobs posted yet</p>

	<a href="mailto:contact@mu.xyz">Post a job</a>
	`)
	w.Write([]byte(html))
}

func Register() {}
