package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rogeriopiatek/goLang-todoAPI/models"
)

/*
HTTP endpoint responsible to call Update.
This will change the Todo based on the URL's ID parameter
*/
func Update(w http.ResponseWriter, r *http.Request) {
	//every URL parameter comes as string, here it's converted to Int
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	//decodes the request body (JSON) to our `todo` variable
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//calls our Update, to change data on the DB
	rows, err := models.Update(id, todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados atualizados com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
