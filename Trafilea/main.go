package main

import (
	"net/http"

	"TRAFILEA/app"

	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

func main() {
	logger := httplog.NewLogger("http-log", httplog.Options{
		Concise: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))

	router(r)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func router(r *chi.Mux) {
	operations := app.DependencyInjection()

	r.Get("/multiples/{num}", operations.GetMultipliers) //all collection
	r.Post("/multiples/add/{num}", operations.SaveNumberCollection)
	r.Get("/multiples/multiplesCollection", operations.GetCollection)
	r.Get("/multiples/value/{num}", operations.GetValueByNumber)

}
