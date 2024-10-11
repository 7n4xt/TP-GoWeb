package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
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

type FormStorage struct {
	IsValid bool
	Value   string
}

type DisplayPage struct {
	IsValid   bool
	LastName  string
	FirstName string
	Date      string
	Gender    string
	IsEmpty   bool
}

var formStorageLastName = FormStorage{false, ""}
var formStorageFirstName = FormStorage{false, ""}
var formStorageDate = FormStorage{false, ""}
var formStorageGender = FormStorage{false, ""}

func main() {
	temp, tempErr := template.ParseGlob("Templates/*.html")
	if tempErr != nil {
		fmt.Printf("Error loading templates: %s", tempErr.Error())
		os.Exit(2)
	}

	changeState := &Change{
		Pair:    false,
		Counter: 0,
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Ynov{
			Title:      "B1 CyberSecurity",
			Sector:     "Cyber Security",
			Level:      "B1",
			NbrStudent: 8,
			Users: []User{
				{FirstName: "Abdulmalek", LastName: "ESUGHI", Age: 20, Sex: true},
				{FirstName: "Enzo", LastName: "ROSSI", Age: 18, Sex: true},
				{FirstName: "leo", LastName: "GOMEZ", Age: 23, Sex: true},
				{FirstName: "Maxime", LastName: "DEBRUN", Age: 17, Sex: true},
				{FirstName: "Adrien", LastName: "DIRIX", Age: 20, Sex: true},
				{FirstName: "Anissa", LastName: "BOUKERCHE", Age: 18, Sex: false},
				{FirstName: "Sylia", LastName: "ABOUD", Age: 18, Sex: false},
				{FirstName: "Eddy", LastName: "AMIR", Age: 18, Sex: true},
			},
		}
		if err := temp.ExecuteTemplate(w, "index", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		changeState.Counter++
		changeState.Pair = changeState.Counter%2 == 0
		if err := temp.ExecuteTemplate(w, "changePage", changeState); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		if err := temp.ExecuteTemplate(w, "Form", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/error?code=400&message=Oops+incorrect+method", http.StatusSeeOther)
			return
		}

		checkValueLastName, _ := regexp.MatchString("^[\\p{L}-]{1,32}$", r.FormValue("lastname"))
		if !checkValueLastName {
			formStorageLastName = FormStorage{false, ""}
			http.Redirect(w, r, "/error?code=400&message=Oops+invalid+last+name+data", http.StatusSeeOther)
			return
		}
		formStorageLastName = FormStorage{true, r.FormValue("lastname")}

		checkValueFirstName, _ := regexp.MatchString("^[\\p{L}-]{1,32}$", r.FormValue("firstname"))
		if !checkValueFirstName {
			formStorageFirstName = FormStorage{false, ""}
			http.Redirect(w, r, "/error?code=400&message=Oops+invalid+first+name+data", http.StatusSeeOther)
			return
		}
		formStorageFirstName = FormStorage{true, r.FormValue("firstname")}

		date := r.FormValue("date")
		if date == "" {
			http.Redirect(w, r, "/error?code=400&message=Date+is+required", http.StatusSeeOther)
			return
		}
		formStorageDate = FormStorage{true, date}

		gender := r.FormValue("gender")
		if gender != "male" && gender != "female" && gender != "other" {
			http.Redirect(w, r, "/error?code=400&message=Invalid+gender+value", http.StatusSeeOther)
			return
		}
		formStorageGender = FormStorage{true, gender}

		http.Redirect(w, r, "/user/display", http.StatusSeeOther)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		data := DisplayPage{
			IsValid:   formStorageLastName.IsValid && formStorageFirstName.IsValid && formStorageDate.IsValid && formStorageGender.IsValid,
			LastName:  formStorageLastName.Value,
			FirstName: formStorageFirstName.Value,
			Date:      formStorageDate.Value,
			Gender:    formStorageGender.Value,
			IsEmpty: !formStorageLastName.IsValid && formStorageLastName.Value == "" &&
				!formStorageFirstName.IsValid && formStorageFirstName.Value == "" &&
				!formStorageDate.IsValid && formStorageDate.Value == "" &&
				!formStorageGender.IsValid && formStorageGender.Value == "",
		}
		if err := temp.ExecuteTemplate(w, "Display", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		message := r.FormValue("message")
		if code != "" && message != "" {
			if err := temp.ExecuteTemplate(w, "Error", map[string]string{
				"Code":    code,
				"Message": message,
			}); err != nil {
				http.Error(w, "Error rendering error page", http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "Sorry, an error occurred", http.StatusInternalServerError)
	})

	fs := http.FileServer(http.Dir("./design"))
	http.Handle("/design/", http.StripPrefix("/design/", fs))

	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Printf("Server error: %s\n", err)
		os.Exit(1)
	}
}
