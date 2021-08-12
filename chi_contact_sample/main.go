package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

type Contact struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	Company     string
	CreatedOn   time.Time
}

var db *gorm.DB
var err error

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", getAllContact)                 //GET /contacts
		r.Get("/{phonenumber}", getContact)       //GET /contacts/0147344454
		r.Post("/", addContact)                   //POST /contacts
		r.Put("/{phonenumber}", updateContact)    //PUT /contacts/0147344454
		r.Delete("/{phonenumber}", deleteContact) //DELETE /contacts/0147344454
	})
	return r
}

func main() {
	host := "localhost"
	dbUser := "postgres"
	dbName := "chi_contact_sample"
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

	db.AutoMigrate(&Contact{})

	// r := chi.NewRouter()
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.Logger)

	// r.Get("/contacts", getAllContact)
	// r.Get("/contacts/{phonenumber}", getContact)
	// r.Post("/contacts", addContact)
	// r.Delete("/contacts/{id}", deleteContact)
	// r.Put("/contacts/{phonenumber}", updateContact)

	r := registerRoutes()
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8002", handler))
}

func getContact(w http.ResponseWriter, r *http.Request) {
	phoneNumber := chi.URLParam(r, "phonenumber")
	fmt.Println(phoneNumber)
	if phoneNumber == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	contact := &Contact{}
	// err := db.First(&contact, phoneNumber)
	err := db.Where("phone_number = ?", phoneNumber).First(&contact)
	if err.Error != nil {
		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s not found", phoneNumber), 404)
		return
	}
	json.NewEncoder(w).Encode(contact)
}

func getAllContact(w http.ResponseWriter, r *http.Request) {
	var contacts []Contact
	db.Find(&contacts)
	json.NewEncoder(w).Encode(&contacts)
}

func addContact(w http.ResponseWriter, r *http.Request) {
	existingContact := &Contact{}
	var contact Contact
	json.NewDecoder(r.Body).Decode(&contact)
	contact.CreatedOn = time.Now()
	// err := db.First(existingContact, contact.PhoneNumber)
	err := db.Where("phone_number = ?", &contact.PhoneNumber).First(&existingContact)
	if err.Error == nil {
		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s already exist", contact.PhoneNumber), 400)
		return
	}
	err = db.Create(&contact)
	if err.Error != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Contact created successfully"))
	w.WriteHeader(201)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	existingContact := &Contact{}
	phoneNumber := chi.URLParam(r, "phonenumber")
	if phoneNumber == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	// err := db.First(existingContact, phoneNumber)
	err := db.Where("phone_number = ?", phoneNumber).First(&existingContact)
	if err.Error != nil {
		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s does not exist", phoneNumber), 400)
		return
	}
	err = db.Delete(&existingContact)
	if err.Error != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Contact deleted"))
	w.WriteHeader(200)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	existingContact := &Contact{}
	phoneNumber := chi.URLParam(r, "phonenumber")
	if phoneNumber == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	var contact Contact
	json.NewDecoder(r.Body).Decode(&contact)
	err := db.Where("phone_number = ?", phoneNumber).First(&existingContact)
	if err.Error != nil {
		http.Error(w, fmt.Sprintf("Contact with phonenumber: %s does not exist", phoneNumber), 400)
		return
	}

	existingContact.FirstName = contact.FirstName
	existingContact.LastName = contact.LastName
	existingContact.Email = contact.Email
	existingContact.Address = contact.Address
	existingContact.Company = contact.Company
	existingContact.PhoneNumber = contact.PhoneNumber

	err = db.Save(&existingContact)

	if err.Error != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Contact update successful"))
	w.WriteHeader(200)
}
