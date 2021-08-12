package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {

	router := chi.NewRouter()

	router.Use(app.JwtAuthentication)
	router.Route("/api", func(router chi.Router) {
		router.Post("/user/signup", controllers.CreateAccount)
		router.Post("/user/login", controllers.Authenticate)
		router.Post("/contacts/new", controllers.CreateContact)
		router.Get("/me/contacts", controllers.GetContactsFor)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
