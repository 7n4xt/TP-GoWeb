package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Sex       bool
}

type Ynov struct {
	Title      string
	Sector     string
	Level      string
	NbrStudent int
	Users      []User
}

type Change struct {
	Pair    bool
	Counter int
}

type StockageForm struct {
	CheckValue bool
	Value      string
}

var stockageForm = StockageForm{false, ""}

func main() {
	temp, tempErr := template.ParseGlob("Templates/*.html")
	if tempErr != nil {
		fmt.Printf("Oups erreur avec le chargement du Template %s", tempErr.Error())
		os.Exit(2)
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Ynov{
			Title:      "B1 Cybersecurity",
			Sector:     "Cyber Security",
			Level:      "B1",
			NbrStudent: 2,
			Users: []User{
				{FirstName: "Abdulmalek", LastName: "ESUGHI", Age: 20, Sex: true},
				{FirstName: "Enzo", LastName: "ROSSI", Age: 18, Sex: true},
			},
		}
		temp.ExecuteTemplate(w, "index", data)
	})

	changeState := &Change{
		Pair:    false,
		Counter: 0,
	}

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		changeState.Counter++
		changeState.Pair = changeState.Counter%2 == 0
		temp.ExecuteTemplate(w, "changePage", changeState)
	})
	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Form", nil)
	})

	fs := http.FileServer(http.Dir("./design"))
	http.Handle("/design/", http.StripPrefix("/design/", fs))

	fmt.Println("Server starting on :8080")
	http.ListenAndServe("localhost:8080", nil)
}
