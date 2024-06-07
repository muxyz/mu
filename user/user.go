// Package user is for authentication
package user

import (
	"encoding/base64"
	"errors"
	"sync"

	"github.com/google/uuid"
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

type Session string

func (s *Session) ID() string {
	return string(*s)
}

func newSess() *Session {
	id := uuid.New().String()
	sess := Session(base64.StdEncoding.EncodeToString([]byte(id)))
	return &sess
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

	sess := newSess()
	sessions[sess.ID()] = sess

	mu.Save(sessions, "sessions.enc", true)

	return acc, sess, nil
}

// Logout the user
func Logout(username string, sess *Session) error {
	mutex.Lock()
	defer mutex.Unlock()

	// logout
	sessions[sess.ID()] = nil

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
		ID:       uuid.New().String(),
		Username: username,
		Password: string(pw),
	}
	mutex.Unlock()

	// save accounts
	return mu.Save(users, "users.enc", true)
}

// Verify a session
func Verify(sess *Session) error {
	mutex.Lock()
	defer mutex.Unlock()

	v, ok := sessions[sess.ID()]
	if !ok {
		return errors.New("invalid session")
	}
	if v == nil {
		return errors.New("expired session")
	}

	return nil
}
