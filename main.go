package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	temp, tempErr := template.ParseFiles("index.html")
	if tempErr != nil {
		fmt.Printf("erreur de templet %s", tempErr.Error())
		os.Exit(02)
	}
}
