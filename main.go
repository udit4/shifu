package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"shifu/rssFetcher"
)

func handler(w http.ResponseWriter, r *http.Request) {

	newsItems, err := rssFetcher.FetchFacebookNews()
	if err != nil {
		log.Printf("Error fetching facebook news: %+v\n", err)
		http.Error(w, "Error fetching news", http.StatusInternalServerError)
		return
	}

	// log.Printf("news items : %+v\n", newsItems)

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Printf("Error loading template : %+v\n", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"NewsItems": newsItems})
	if err != nil {
		log.Printf("Error executing template : %+v\n", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
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
