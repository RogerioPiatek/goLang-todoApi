package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rogeriopiatek/goLang-todoAPI/configs"
	"github.com/rogeriopiatek/goLang-todoAPI/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	//r == router
	r := chi.NewRouter()
	//we don't call the function, but point to it
	r.Post("/", handlers.Create)
	//this syntax retrieves and passes to our function an URL parameter
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
