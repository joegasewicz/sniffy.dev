package routes

import (
	"context"
	"encoding/json"
	"github.com/joegasewicz/sniffy.dev/api/models"
	"github.com/joegasewicz/sniffy.dev/api/schemas"
	"github.com/joegasewicz/sniffy.dev/api/utils"
	"log"
	"net/http"
	"strconv"
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
			Relation("Paths").
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
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var path schemas.Path
		err := decoder.Decode(&path)
		if err != nil {
			return
		}
		idInt, err := strconv.Atoi(id)
		if err != nil || idInt == 0 {
			return
		}
		path.DomainID = int64(idInt)
		// TODO check if domain exists
		ctx := context.Background()
		_, err = utils.Database().
			NewInsert().
			Model(&path).
			Exec(ctx)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errorMsg := schemas.ErrorMessage{Message: "error storing new path"}
			json.NewEncoder(w).Encode(errorMsg)
			return
		}
		log.Printf("new path created: %v\n", path.Name)
		json.NewEncoder(w).Encode(path)
		return

	}
}
