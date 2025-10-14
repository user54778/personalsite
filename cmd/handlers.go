package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	currTime := time.Now().In(loc)

	tmplData := struct {
		Time string
	}{
		Time: currTime.Format("3:04 PM MST"),
	}

	// NOTE: $HOME env not expanded automatically in Go strings
	files := []string{
		os.ExpandEnv("$HOME/code/personalsite/ui/html/pages/base.tmpl"),
		os.ExpandEnv("$HOME/code/personalsite/ui/html/pages/home.tmpl"),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// use ExecuteTemplate() to write content of `base` template as response body
	err = ts.ExecuteTemplate(w, "base", tmplData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("About\n"))
}
