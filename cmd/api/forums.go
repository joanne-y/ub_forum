// Filename: cmd/api/forums.go
package main

import (
	"fmt"
	"net/http"
)

// createForumHandler for the "POST /v1/forums" endpoint
func (app *application) createForumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new forum..")
}

// showForumHandler for the "GET /v1/forums/:id" endpoint
func (app *application) showForumHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Display the forum id
	fmt.Fprintf(w, "show the details for forum %d\n", id)
}
