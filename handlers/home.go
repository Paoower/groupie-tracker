package handlers

import (
	"errors"
	"groupie-tracker/utils"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("%v Tried to access unexistant route %v%v\n", utils.YELLOW, r.URL.Path, utils.NONE)
		ErrorHandler(w, errors.New("404 | Not found"))
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, errors.New("400 | Wrong method"))
		return
	}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, errors.New("500 | Internal server error: cannot parse template"))
		return
	}

	t.Execute(w, nil)
}
