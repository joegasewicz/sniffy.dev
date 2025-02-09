package routes

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.gohtml",
		"templates/partials/navbar.gohtml",
		"templates/partials/footer.gohtml",
		"templates/home.gohtml",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", nil)
}
