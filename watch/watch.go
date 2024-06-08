package watch

import (
	"context"
	"fmt"
	"mu.dev"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var Key = os.Getenv("YOUTUBE_API_KEY")
var Client, _ = youtube.NewService(context.TODO(), option.WithAPIKey(Key))

func watchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// search request
		r.ParseForm()
		q := r.Form.Get("q")

		resp, err := Client.Search.List([]string{"id", "snippet"}).Q(q).MaxResults(25).Do()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var results string

		for _, item := range resp.Items {
			var id, url, desc string
			kind := strings.Split(item.Id.Kind, "#")[1]
			t, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			desc = fmt.Sprintf(`[%s] published on %s`, kind, t.Format(time.RFC822))
			switch kind {
			case "video":
				id = item.Id.VideoId
				url = "https://www.youtube.com/watch?v=" + id
			case "playlist":
				id = item.Id.PlaylistId
				url = "https://www.youtube.com/playlist?list=" + id
			case "channel":
				id = item.Id.ChannelId
				url = "https://www.youtube.com/channel/" + id
				desc = "[channel]"
			}
			channel := fmt.Sprintf(`<a href="https://youtube.com/channel/%s">%s</a>`, item.Snippet.ChannelId, item.Snippet.ChannelTitle)
			results += fmt.Sprintf(`
				<div class="video"><a href="%s"><img src="%s"><h3>%s</h3></a>%s | %s</div>`,
				url, item.Snippet.Thumbnails.Medium.Url, item.Snippet.Title, channel, desc)
		}

		html := mu.Template("Watch", "Results", "", fmt.Sprintf(`
<style>
  form {
    margin-top: 100px;
  }
  .video {
    margin-bottom: 50px;
  }
  img {
    border-radius: 10px;
  }
  h3 {
    margin-bottom: 5px;
  }
</style>
<form action="/watch" method="POST">
  <input name="q" id="q" placeholder=Search>
  <button>Submit</button>
</form>
<h1>Results</h1>
<div id="results">
%s
</div>`, results))
		w.Write([]byte(html))
		return
	}

	html := mu.Template("Watch", "Watch YouTube Videos", "", `
<style>
form {
  margin-top: 100px;
}
</style>
<form action="/watch" method="POST">
  <input name="q" id="q" placeholder=Search>
  <button>Submit</button>
</form>`)

	w.Write([]byte(html))
}

func Register() {
	http.HandleFunc("/watch", watchHandler)
}
