package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// var router *chi.Mux

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "id")
	// For Gorilla MUX router
	// vars := mux.Vars(r)
	// key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func handleRequests() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Get("/", homePage)
		r.Get("/articles", returnAllArticles)
		r.Get("/article/{id}", returnSingleArticle)
		r.Post("/article", createNewArticle)
		r.Delete("/article/{id}", deleteArticle)
	})

	// USING GORILLA MUX ROUTER
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
