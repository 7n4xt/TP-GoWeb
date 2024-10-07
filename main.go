package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	temp, tempErr := template.ParseFiles("*.html")
	if tempErr != nil {
		fmt.Printf("erreur de templet %s", tempErr.Error())
		os.Exit(02)
	}

	type User struct {
		FristName string
		LastName  string
	}
	type Ynov struct {
		titre       string
		Filiere     string
		Niveau      string
		NbrEtudiant int
		Users       []User
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Ynov{"307", "Cyber", "B1", 5, []User{{"Abdulmalek", "ESUGHI"}, {"Enzo", "ROSSI"}}}
		temp.ExecuteTemplate(w, "index", data)
	})
}
