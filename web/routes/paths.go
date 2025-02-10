package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	form_validator "github.com/joegasewicz/form-validator"
	"github.com/joegasewicz/sniffy.dev/web/schemas"
	"github.com/joegasewicz/sniffy.dev/web/utils"
	"io"
	"log"
	"net/http"
	"strconv"
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
	if r.Method == "GET" {
		domains := []schemas.Domain{}
		GetDomainData(w, &domains)

		id := r.PathValue("id")
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		domain := schemas.Domain{}
		url := fmt.Sprintf("http://localhost:3000/domains/%s/paths", id)
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
		return
	}
	if r.Method == "POST" {
		id := r.PathValue("id")
		ideInt, err := strconv.Atoi(id)
		if err != nil {
			return
		}
		c := form_validator.Config{
			MaxMemory: 0,
			Fields: []form_validator.Field{
				{
					Name:     "path",
					Validate: true,
					Type:     "string",
				},
			},
		}
		if ok := form_validator.ValidateForm(r, &c); ok {
			name, _ := form_validator.GetString("path", &c)
			var path schemas.Path
			path.Name = name
			path.DomainID = uint64(ideInt)
			jsonBytes, _ := json.Marshal(&path)
			req, err := http.NewRequest(
				"POST",
				fmt.Sprintf("http://localhost:3000/domains/%d/paths", ideInt),
				bytes.NewBuffer(jsonBytes),
			)
			if err != nil {
				log.Printf("error marshalling JOSN\n")
				w.WriteHeader(http.StatusInternalServerError)
				utils.SetTemplate(w, "templates/paths.gohtml", nil)
				return
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			defer resp.Body.Close()
			log.Printf("successfully sent POST to /domains/{id}/paths")
			http.Redirect(w, r, "/domain/"+id+"", http.StatusSeeOther)
		} else {
			var formErrs = form_validator.FormErrors{}
			form_validator.GetFormErrors(&c, &formErrs)
			log.Printf("form errors: %s", formErrs)
			http.Redirect(w, r, "/domains/"+id+"/paths", http.StatusSeeOther)
		}
	}

}
