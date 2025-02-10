package routes

import (
	"context"
	"encoding/json"
	"github.com/joegasewicz/sniffy.dev/api/models"
	"github.com/joegasewicz/sniffy.dev/api/schemas"
	"github.com/joegasewicz/sniffy.dev/api/utils"
	"log"
	"net/http"
)

func Domain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		ctx := context.Background()
		var domain []models.Domain
		err := utils.Database().
			NewSelect().
			Model(&domain).
			Relation("Paths").
			Scan(ctx)
		if err != nil {
			log.Printf("error retrieving domains\n")
			w.WriteHeader(http.StatusInternalServerError)
			errMsg := schemas.ErrorMessage{Message: "error retrieving domains"}
			json.NewEncoder(w).Encode(errMsg)
			return
		}
		json.NewEncoder(w).Encode(domain)
	}
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var domain schemas.Domain
		err := decoder.Decode(&domain)
		if err != nil {
			return
		}
		ctx := context.Background()
		_, err = utils.Database().
			NewInsert().
			Model(&domain).
			Exec(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorMsg := schemas.ErrorMessage{Message: "error storing domain to database."}
			json.NewEncoder(w).Encode(errorMsg)
			return
		}
		log.Printf("new domain created: %v\n", domain.Name)
		json.NewEncoder(w).Encode(domain)
		return
	}
}
