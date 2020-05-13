package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Print(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Print(err)
	}
}
