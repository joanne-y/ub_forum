// Filename: cmd/api/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Create a new httprouter router instance
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/forums", app.listForumsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/forums", app.createForumHandler)
	router.HandlerFunc(http.MethodGet, "/v1/forums/:id", app.showForumHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/forums/:id", app.updateForumHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/forums/:id", app.deleteForumHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
