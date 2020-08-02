package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Look for a user present on the dbUsers data structure.
// If the client do not have a cookie so far, creates one.
func getUser(w http.ResponseWriter, req *http.Request) User {
	cookie, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewRandom()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	var u User
	if session, ok := dbSessions[cookie.Value]; ok {
		session.lastActivity = time.Now()
		dbSessions[cookie.Value] = session
		u = dbUsers[session.un]
	}

	return u
}

// Verify if a user is already loggedIn
func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	session, ok := dbSessions[cookie.Value]
	if ok {
		session.lastActivity = time.Now()
		dbSessions[cookie.Value] = session
	}

	_, ok = dbUsers[session.un]
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	return ok
}

// Clean all the sessions that have expired.
func cleanSessions() {
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * time.Duration(sessionLength)) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
}
