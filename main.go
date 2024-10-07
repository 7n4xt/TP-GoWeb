package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	temp, tempErr := template.ParseFiles("index.html")
	if tempErr != nil {
		fmt.Printf("erreur de templet %s", tempErr.Error())
		os.Exit(02)
	}
	type Ynov struct {
		titre       string
		filiere     string
		niveau      string
		nbrEtudiant int
	}
	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {

	})
}
