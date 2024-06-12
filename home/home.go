package home

import (
	"net/http"

	"mu.dev"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	html := mu.Template("Home", "Home screen", `<a href="/logout">Logout</a>`, `
<style>
#title {
  margin-top: 100px;
}
.apps a {
  margin-right: 10px;
}
</style>
          <h1 id="title">Home</h1>
          <p id="description"></p>

	  <div class="apps">
	    <a href="/chat">
	      <button>
		Chat
	      </button>
	    </a>
	    <a href="/news">
	      <button>
		News
	      </button>
	    </a>
	    <a href="/pray">
	      <button>
		Pray
	      </button>
	    </a>
	    <a href="/reminder">
	      <button>
		Reminder
	      </button>
	    </a>
	    <a href="/watch">
	      <button>
		Watch
	      </button>
	    </a>
	  </div>
	`)

	w.Write([]byte(html))
}

func Register() {}
