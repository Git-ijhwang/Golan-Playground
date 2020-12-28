package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user

	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	if un, ok := dbSessions[c.Value]; ok {
		un.lastActivity = time.Now()
		dbSessions[c.Value] = un
		u = dbUsers[un.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un, ok := dbSessions[c.Value]
	if ok {
		un.lastActivity = time.Now()
		dbSessions[c.Value] = un
	}

	_, ok = dbUsers[un.un]

	//refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	return ok
}

func cleanSessions() {
	fmt.Println("Before Clean")
	showSessions()

	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}

	dbSessionsCleaned = time.Now()

	fmt.Println("After Clean")
	showSessions()
}

func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
