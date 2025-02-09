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

func PathHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		ctx := context.Background()
		var domain models.Domain
		err := utils.Database().
			NewSelect().
			Model(&domain).
			Where("id = ?", id).
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
}
