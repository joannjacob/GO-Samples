package main

import (
	"fmt"
	"net/http"
	"os"

	"chi_post_sample/driver"
	ph "chi_post_sample/handler/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	dbName := "chi_post_sample"
	dbUser := "postgres"
	dbPass := "postgres"
	dbHost := "localhost"
	dbPort := "5432"

	println("DB values", dbName, dbHost, dbPass, dbPort, dbUser)

	connection, err := driver.ConnectSQL(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	// pHandler := cors.Default().Handler(r)
	r.Get("/posts", pHandler.Fetch)
	r.Get("/posts/{id}", pHandler.GetByID)
	r.Post("/posts/create", pHandler.Create)
	r.Put("/posts/update/{id}", pHandler.Update)
	r.Delete("/posts/{id}", pHandler.Delete)

	fmt.Println("Server listening at :8005")
	http.ListenAndServe(":8005", r)
}
