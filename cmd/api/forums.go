// Filename: cmd/api/forums.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"ubforum.joanneyong.net/internal/data"
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
	// Create a new instance of the Forum struct containing the ID we extracted
	// from our URL and some sample data
	forum := data.Forum{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "UB nursing program",
		Message:   "How is the nursing program at UB?",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, forum, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
