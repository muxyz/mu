package work

import (
	"net/http"

	"mu.dev"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	html := mu.Template("Work", "Do good work", "", `
	<h1 style="padding-top: 100px;">Work</h1>
	<h3>Find work, do work, get paid!</h3>
	<p>No jobs posted yet</p>

	<a href="mailto:contact@mu.xyz">Post a job</a>
	`)
	w.Write([]byte(html))
}

func Register() {}
