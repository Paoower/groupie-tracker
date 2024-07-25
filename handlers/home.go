package handlers

import (
	"errors"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, errors.New("400 | Wrong method"))
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: cannot parse template"))
	}

	t.Execute(w, nil)
}
