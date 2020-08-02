package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	user = iota
	admin
	agent
)

type User struct {
	Username  string
	Password  []byte
	FirstName string
	LastName  string
	Role      int
}

type Session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]User{}
var dbSessions = map[string]Session{}
var dbSessionsCleaned time.Time

var sessionLength = 30

func init() {
	// Parse all the templates present on the template folder.
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
	// Applies the authorized middleware.
	http.HandleFunc("/logout", authorized(logout))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Welcome page, shows the user name and bar link if the current client is
// logged in, if is not, shows the login and signup links.
func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)

	err := tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		_, _ = io.WriteString(w, "Resource not valid")
		http.Error(w, "", http.StatusBadRequest)
		log.Fatalln(err)
	}
	fmt.Println(user, admin)
}

// Redirect outside if the current user not have the permission set as Agent.
func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != agent {
		http.Error(w, "You must be an agent to enter the bar", http.StatusForbidden)
		return
	}

	_ = tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// Creates an user, encrypt a password and creates the cookie and session.
func signUp(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var user User
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		role, _ := strconv.Atoi(req.FormValue("role"))
		if _, ok := dbUsers[username]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		sID, _ := uuid.NewRandom()
		cookie := &http.Cookie{
			Name:   "session",
			Value:  sID.String(),
			MaxAge: sessionLength,
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = Session{username, time.Now()}
		bc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		user = User{
			Username:  username,
			Password:  bc,
			FirstName: req.FormValue("first-name"),
			LastName:  req.FormValue("last-name"),
			Role:      role,
		}
		dbUsers[user.Username] = user
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	_ = tpl.ExecuteTemplate(w, "signup.gohtml", user)
}

// Makes a relation between a session present on the cookie, and a username, then
// verifies if the password correspond to the stored one.
func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var user User
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		// Verifies if the username even exists.
		u, ok := dbUsers[username]
		if !ok {
			http.Error(w, "Username or password does not match", http.StatusForbidden)
			return
		}
		// Compare hash with the given password.
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil {
			http.Error(w, "Username or password does not match", http.StatusForbidden)
			return
		}
		// Creates a new cookie and a new session.
		sID, _ := uuid.NewRandom()
		c := &http.Cookie{
			Name:   "session",
			Value:  sID.String(),
			MaxAge: sessionLength,
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = Session{username, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	_ = tpl.ExecuteTemplate(w, "login.gohtml", user)
}

// Set the time of the cookie to -1, that means session expired
func logout(w http.ResponseWriter, req *http.Request) {
	c, _ := req.Cookie("session")
	delete(dbSessions, c.Value)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

// Creates a middleware to verify if a user is already logged in or if should
// redirects.
func authorized(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if !alreadyLoggedIn(w, req) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, req)
	}
}
