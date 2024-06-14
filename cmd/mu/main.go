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
	"mu.dev/watch"
)

func main() {
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`User-agent: *
Allow: /`))
	})

	// chat
	http.HandleFunc("/chat", user.Auth(chat.IndexHandler))
	http.HandleFunc("/chat/prompt", user.Auth(chat.PromptHandler))
	http.HandleFunc("/chat/channels", user.Auth(chat.ChannelHandler))

	// home
	http.HandleFunc("/home", user.Auth(home.IndexHandler))

	// news
	http.HandleFunc("/news", user.Auth(news.IndexHandler))
	http.HandleFunc("/news/feeds", user.Auth(news.FeedsHandler))
	http.HandleFunc("/news/status", user.Auth(news.StatusHandler))
	// http.HandleFunc("/add", addHandler)

	// pray
	http.HandleFunc("/pray", user.Auth(pray.IndexHandler))

	// reminder
	http.HandleFunc("/reminder", user.Auth(reminder.IndexHandler))

	// user
	http.HandleFunc("/login", user.LoginHandler)
	http.HandleFunc("/logout", user.LogoutHandler)
	http.HandleFunc("/signup", user.SignupHandler)

	// user admin
	http.HandleFunc("/admin", user.Auth(user.Admin))

	// watch
	http.HandleFunc("/watch", user.Auth(watch.WatchHandler))

	// work
	http.HandleFunc("/work", work.Handler)

	// any other stuff
	chat.Register()
	home.Register()
	news.Register()
	pray.Register()
	reminder.Register()
	user.Register()
	watch.Register()
	work.Register()

	mu.Serve(8080)
}
