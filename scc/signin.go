package scc

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var (
	tmpl = template.Must(template.ParseFiles(
		"tmpl/signin.html",
	))
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(os.Getwd())
	err := tmpl.ExecuteTemplate(w, "signin.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
