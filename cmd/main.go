package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	/*
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>Hello, World!</h1>"))
		})
	*/
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))

	log.Println("Starting server on :8080")
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
