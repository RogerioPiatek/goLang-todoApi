package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/rogeriopiatek/goLang-todoAPI/models"
)

/*
HTTP endpoint responsible to call Update.
This will change the Todo based on the URL's ID parameter
*/
func Delete(w http.ResponseWriter, r *http.Request) {
	//every URL parameter comes as string, here it's converted to Int
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//calls our Update, to change data on the DB
	rows, err := models.Delete(id)
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Registro removido com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
