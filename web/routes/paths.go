package routes

import (
	"encoding/json"
	"fmt"
	"github.com/joegasewicz/sniffy.dev/web/schemas"
	"github.com/joegasewicz/sniffy.dev/web/utils"
	"io"
	"log"
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		domains := []schemas.Domain{}
		resp, err := http.Get("http://localhost:3000/domains")
		if err != nil {
			log.Printf("error requesting domains")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &domains); err != nil {
			log.Printf("error unmarshalling domains")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		templateData := map[string][]schemas.Domain{"Domains": domains}
		utils.SetTemplate(w, "templates/paths.gohtml", templateData)

	}
}

func DomainPathsHandler(w http.ResponseWriter, r *http.Request) {
	domains := []schemas.Domain{}
	GetDomainData(w, &domains)

	id := r.PathValue("id")
	fmt.Print("here-------> ", id)
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	domain := schemas.Domain{}
	url := fmt.Sprintf("http://localhost:3000/domains/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error requesting domains")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &domain); err != nil {
		log.Printf("error unmarshalling domains")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	templateData := map[string]interface{}{"Domain": domain, "Domains": domains}
	utils.SetTemplate(w, "templates/paths.gohtml", templateData)

}
