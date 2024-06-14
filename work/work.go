package work

import (
	"net/http"

	"mu.dev"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	nav := `<a href="mailto:contact@mu.xyz" class="head">Post a job</a>`
	html := mu.Template("Work", "Do work, get paid", nav, `
	<h1 style="padding-top: 100px">Work</h1>
	<p>No jobs posted yet</p>
	`)
	mu.Render(w, html)
}

func Register() {}
