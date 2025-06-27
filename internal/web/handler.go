package web

import (
	"html/template"
	"log"
	"net/http"
	"tecalliance-link/internal/config"
)

type pageData struct {
	Title string
	Links []config.Link
}

func LinksHandler() http.HandlerFunc {

	tmpl, err := template.ParseFiles("internal/web/templates/index.gohtml")
	if err != nil {
		log.Fatalf("Error: failed to load page template: %v", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {

		data := pageData{
			Title: "TecAlliance Links",
			Links: config.Links,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
