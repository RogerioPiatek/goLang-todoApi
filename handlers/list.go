package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rogeriopiatek/goLang-todoAPI/models"
)

// HTTP endpoint that will call GetAll and return the JSON of Todos
func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros")
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
