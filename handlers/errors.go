package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func ErrorHandler(w http.ResponseWriter, err error) {
	// Parse the error message
	parts := strings.SplitN(err.Error(), "|", 2)
	statusStr := strings.TrimSpace(parts[0])
	errMsg := strings.TrimSpace(parts[1])

	// Parse the status code
	var statusCode int
	_, err = fmt.Sscanf(statusStr, "%d", &statusCode)
	if err != nil {
		// If we can't parse the status code, default to 500
		statusCode = http.StatusInternalServerError
	}

	// Prepare the data for the template
	data := struct {
		ErrStatus string
		ErrMsg    string
	}{
		ErrStatus: statusStr,
		ErrMsg:    errMsg,
	}

	// Parse the template
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// If we can't parse the template, fall back to a simple error message
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the status code
	w.WriteHeader(statusCode)

	// Execute the template
	if err := t.Execute(w, data); err != nil {
		// If we can't execute the template, log the error
		// We can't write an error response here because headers have already been sent
		log.Printf("Error executing template: %v", err)
	}
}
