package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Printf("Error loading template : %+v\n", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	if _, err := os.Stat("template.html"); os.IsNotExist(err) {
		log.Fatal("template.html file not found")
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	log.Println("Starting server on :9000")
	http.ListenAndServe(":9000", nil)
}
