package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error

type Resource struct {
	gorm.Model

	Link        string
	Name        string
	Author      string
	Description string
	Tags        pq.StringArray `gorm:"type:varchar(64)[]"`
}

func GetResources(w http.ResponseWriter, r *http.Request) {
	var resources []Resource
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}
func GetResource(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	var resource Resource
	db.First(&resource, params)
	json.NewEncoder(w).Encode(&resource)
}
func CreateResource(w http.ResponseWriter, r *http.Request) {
	var resource Resource
	json.NewDecoder(r.Body).Decode(&resource)
	db.Create(&resource)
	json.NewEncoder(w).Encode(&resource)
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	var resource Resource
	db.First(&resource, params)
	db.Delete(&resource)

	var resources []Resource
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	host := "localhost"
	dbUser := "postgres"
	dbName := "gorm_sample"
	password := "postgres"

	db, err = gorm.Open(
		"postgres",
		"host="+host+" user="+dbUser+
			" dbname="+dbName+" sslmode=disable password="+
			password)

	if err != nil {
		fmt.Println("DB connection failed")
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Resource{})

	r.Get("/resources", GetResources)
	r.Get("/resources/{id}", GetResource)
	r.Post("/resources", CreateResource)
	r.Delete("/resources/{id}", DeleteResource)

	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
