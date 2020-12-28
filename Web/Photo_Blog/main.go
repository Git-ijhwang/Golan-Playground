package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./Public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)

		/* Create New file */
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}

		defer nf.Close()

		mf.Seek(0, 0)
		io.Copy(nf, mf)

		//add filename to this use
		c = appendValue(w, c, fname)

	}
	//c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}

	c.Value = s
	http.SetCookie(w, c)
	return c

	/*
		p1 := "disneyland.jpg"
		p2 := "atbeach.jpg"
		p3 := "hollywood.jpg"

		s := c.Value
		if !strings.Contains(s, p1) {
			s += "|" + p1
		}
		if !strings.Contains(s, p2) {
			s += "|" + p2
		}
		if !strings.Contains(s, p3) {
			s += "|" + p3
		}

		c.Value = s
		http.SetCookie(w, c)
		return c
	*/
}
