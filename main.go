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
	titre      string
	Sector     string
	Level      string
	NbrStudent int
	Users      []User
}
type Change struct {
	Pair    bool
	Counter int
}

func main() {
	temp, tempErr := template.ParseGlob("Templates/*.html")
	if tempErr != nil {
		fmt.Printf("Oups erreur avec le chargement du Template %s", tempErr.Error())
		os.Exit(02)
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Ynov{"307", "Cyber", "B1", 2, []User{{"Abdulmalek", "ESUGHI", 20, "male"}, {"Enzo", "ROSSI", 18, "male"}}}
		temp.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe("localhost:8080", nil)
}
