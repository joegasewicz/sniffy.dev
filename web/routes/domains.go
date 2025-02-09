package routes

import (
	"bytes"
	"encoding/json"
	form_validator "github.com/joegasewicz/form-validator"
	"github.com/joegasewicz/sniffy.dev/web/schemas"
	"github.com/joegasewicz/sniffy.dev/web/utils"
	"io"
	"log"
	"net/http"
)

func GetDomainData(w http.ResponseWriter, domains *[]schemas.Domain) {

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
}

func DomainHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		domains := []schemas.Domain{}
		GetDomainData(w, &domains)
		templateData := map[string][]schemas.Domain{"Domains": domains}
		utils.SetTemplate(w, "templates/domains.gohtml", templateData)
	}
	if r.Method == "POST" {
		c := form_validator.Config{
			MaxMemory: 0,
			Fields: []form_validator.Field{
				{
					Name:     "name",
					Validate: true,
					Type:     "string",
				},
			},
		}
		if ok := form_validator.ValidateForm(r, &c); ok {
			name, _ := form_validator.GetString("name", &c)
			var domain schemas.Domain
			domain.Name = name
			jsonBytes, _ := json.Marshal(&domain)
			req, err := http.NewRequest(
				"POST",
				"http://localhost:3000/domains",
				bytes.NewBuffer(jsonBytes),
			)
			if err != nil {
				log.Printf("error marshalling JOSN\n")
				w.WriteHeader(http.StatusInternalServerError)
				utils.SetTemplate(w, "templates/domains.gohtml", nil)
				return
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			defer resp.Body.Close()
			log.Printf("status: %d\n", resp.Status)
			log.Printf("successfully sent POST request")
			http.Redirect(w, r, "/domains", http.StatusSeeOther)
			return
		} else {
			var formErrs = form_validator.FormErrors{}
			form_validator.GetFormErrors(&c, &formErrs)
			log.Printf("form errors: %s", formErrs)
			http.Redirect(w, r, "/domains", http.StatusSeeOther)
		}
	}
}
