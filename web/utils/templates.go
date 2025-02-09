package utils

import (
	templates2 "github.com/joegasewicz/sniffy.dev/web/templates"
	"html/template"
	"log"
	"net/http"
)

func SetTemplate(w http.ResponseWriter, templatePath string, data any) {
	files := append(templates2.TemplateFiles, templatePath)
	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("error loading templates")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
