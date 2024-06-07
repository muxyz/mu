// Package user is for authentication
package user

import (
	"encoding/base64"
	"errors"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"mu.dev"
)

var mutex sync.Mutex
var users = map[string]*Account{}
var sessions = map[string]*Session{}

func init() {
	// load users
	mu.Load(&users, "users.enc", true)
	mu.Load(&sessions, "sessions.enc", true)
}

type Account struct {
	ID       string
	Username string
	Password string
}

type Session struct {
	ID       string
	Username string
}

func newSess(acc *Account) *Session {
	id := mu.ID()
	sess := &Session{
		ID:       base64.StdEncoding.EncodeToString([]byte(id)),
		Username: acc.Username,
	}
	return sess
}

// Login a user
func Login(username, password string) (*Account, *Session, error) {
	mutex.Lock()
	defer mutex.Unlock()

	acc, ok := users[username]
	if !ok {
		return nil, nil, errors.New("does not exist")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)); err != nil {
		return nil, nil, errors.New("invalid password")
	}

	sess := newSess(acc)
	sessions[sess.ID] = sess

	mu.Save(sessions, "sessions.enc", true)

	return acc, sess, nil
}

// Logout the user
func Logout(username string, sess *Session) error {
	mutex.Lock()
	defer mutex.Unlock()

	// logout
	sessions[sess.ID] = nil

	// TODO: ...
	return mu.Save(sessions, "sessions.enc", true)
}

// Signup a user
func Signup(username, password string) error {
	if _, ok := users[username]; ok {
		return errors.New("already exists")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	mutex.Lock()
	users[username] = &Account{
		ID:       mu.ID(),
		Username: username,
		Password: string(pw),
	}
	mutex.Unlock()

	// save accounts
	return mu.Save(users, "users.enc", true)
}

// Verify a session
func Verify(sessID, user string) error {
	mutex.Lock()
	defer mutex.Unlock()

	sess, ok := sessions[sessID]
	if !ok {
		return errors.New("invalid session")
	}
	if sess == nil {
		return errors.New("expired session")
	}
	if user != sess.Username {
		return errors.New("invalid user")
	}

	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		user := r.Form.Get("username")
		pass := r.Form.Get("password")

		acc, sess, err := Login(user, pass)
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "sess",
			Value: sess.ID,
		})

		http.SetCookie(w, &http.Cookie{
			Name:  "user",
			Value: acc.Username,
		})

		http.Redirect(w, r, "/home", 302)
		return
	}

	// Login screen
	html := mu.Template("Login", "Login to your account", "", `
<style>
  #login {
    padding-top: 100px;
  }
</style>
<div id="login">
<h1>Login</h1>
<form action="/login", method="post">
  <input id="username" name="username" placeholder=Username>
  <br><br>
  <input id="password" name="password" type="password" placeholder=Password>
  <br><br>
  <button>Submit</button>
</form>
</div>
`)
	w.Write([]byte(html))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		user := r.Form.Get("username")
		pass := r.Form.Get("password")

		err := Signup(user, pass)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		acc, sess, err := Login(user, pass)
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "sess",
			Value: sess.ID,
		})

		http.SetCookie(w, &http.Cookie{
			Name:  "user",
			Value: acc.Username,
		})

		http.Redirect(w, r, "/home", 302)
		return
	}

	// Login screen
	html := mu.Template("Signup", "Signup for an account", "", `
<style>
  #signup {
    padding-top: 100px;
  }
</style>
<div id="signup">
<h1>Signup</h1>
<form action="/signup", method="post">
  <input id="username" name="username" placeholder=Username>
  <br><br>
  <input id="password" name="password" type="password" placeholder=Password>
  <br><br>
  <button>Submit</button>
</form>
</div>
`)
	w.Write([]byte(html))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "sess",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "user",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", 302)
}

func Register() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/signup", signupHandler)
}
