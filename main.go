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
	Sex       string
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
			Title:      "307",
			Sector:     "Cyber Security",
			Level:      "B1",
			NbrStudent: 2,
			Users: []User{
				{FirstName: "Abdulmalek", LastName: "ESUGHI", Age: 20, Sex: "male"},
				{FirstName: "Enzo", LastName: "ROSSI", Age: 18, Sex: "male"},
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

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorrecte", http.StatusMovedPermanently)
			return
		}

		fs := http.FileServer(http.Dir("./design"))
		http.Handle("/design/", http.StripPrefix("/design/", fs))

		fmt.Println("Server starting on :8080")
		http.ListenAndServe("localhost:8080", nil)
	})
}
