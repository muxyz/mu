package main

import (
	"net/http"

	"mu.dev"
	"mu.dev/chat"
	"mu.dev/home"
	"mu.dev/news"
	"mu.dev/pray"
	"mu.dev/reminder"
	"mu.dev/user"
)

func main() {
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`User-agent: *
Allow: /`))
	})

	chat.Register()
	home.Register()
	news.Register()
	pray.Register()
	reminder.Register()
	user.Register()

	mu.Serve(8080)
}
