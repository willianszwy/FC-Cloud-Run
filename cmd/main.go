package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"willianszwy/FC-Cloud-Run/configs"
	"willianszwy/FC-Cloud-Run/internal/handlers"
)

func main() {

	config, err := configs.LoadConfig("../")
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/temperature", handlers.GetTemperature(config))

	http.ListenAndServe(":8080", r)
}
