package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rogeriopiatek/goLang-todoAPI/models"
)

// An HTTP endpoint that will call the Insert model
func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	//decodes the request body (JSON) to our `todo` variable
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		/*
			The method below, `http.Error`, sends an HTTP error response. It receives three parameters:

			1. `w`: An `http.ResponseWriter` used to construct the HTTP response.
			2. `error`: A string representing the error message, in this case, it uses `http.StatusText` to get the standard text for the HTTP status code.
			3. `code`: The HTTP status code, in this case, `http.StatusInternalServerError` which corresponds to 500 - Internal Server Error.

			This results in an HTTP response with a 500 status code and the associated error message.
		*/
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//calls the Insert defined on the models package, and receives the ID returned by it
	id, err := models.Insert(todo)

	//Creating our JSON response (map of string to any)
	var resp map[string]any

	// If there's an error, prepare a JSON response indicating the error.
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
