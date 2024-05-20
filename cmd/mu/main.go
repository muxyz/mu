package main

import (
	"net/http"

	"mu.dev"
	"mu.dev/apps/chat"
	"mu.dev/apps/news"
	"mu.dev/apps/pray"
	"mu.dev/apps/reminder"
)

func main() {
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`User-agent: *
Allow: /`))
	})

	chat.Register()
	news.Register()
	pray.Register()
	reminder.Register()

	mu.Serve(8080)
}
