package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index (w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func login (w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)

}

func sign_up (w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/sign_up.html"))
	tmpl.Execute(w, nil)
}

func main() {
	fmt.Println("http://localhost:8000")
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/sign_up", sign_up)

	http.ListenAndServe(":8000", nil)
}
