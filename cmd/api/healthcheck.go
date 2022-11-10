// Filename: cmd/api/healthcheck.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: avaialble")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
	// Create a map to hold our healthcheck data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	// Convert our map into a JSON object
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	// Add a newline to make viewing on the terminal easier
	js = append(js, '\n')
	// Specifiy that we will serve our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the []byte slice containing the JSON response body
	w.Write(js)
}
