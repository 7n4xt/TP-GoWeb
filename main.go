package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type User struct {
	FristName string
	LastName  string
	age       int
	sex       string
}
type Ynov struct {
	titre       string
	Filiere     string
	Niveau      string
	NbrEtudiant int
	Users       []User
}

func main() {
	temp, tempErr := template.ParseFiles("index.html")
	if tempErr != nil {
		fmt.Printf("erreur de templet %s", tempErr.Error())
		os.Exit(02)
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Ynov{"307", "Cyber", "B1", 5, []User{{"Abdulmalek", "ESUGHI", 20, "Masculin"}, {"Enzo", "ROSSI", 18, "Masculin"}}}
		temp.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe("localhost:8080", nil)
}
