package main

import (
	"fmt"
	"net/http"
)

// logError logs an error along with some key request details like the method and URL.
// This helps in tracing where and how the error occurred.
func (app *application) logError(r *http.Request, err error) {
	app.logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

// errorResponse is a helper function that sends a JSON-formatted error response to the client.
// It encapsulates the error message and ensures proper logging if something goes wrong during this process.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	// Attempt to send the error response as JSON.
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		// If an error occurs while sending the response, log it and return an internal server error.
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// serverErrorResponse is a specific response for internal server errors (500).
// It logs the original error, sends a generic message to the client, and ensures the client isn't exposed to sensitive error details.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse handles the case where a requested resource is not found (404).
// This function ensures a clean, user-friendly response is sent.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse informs the client when the HTTP method used is not supported for the requested resource (405).
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// badRequestResponse is used when the server receives a malformed or invalid request (400).
// It sends the specific error back to the client.
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// failedValidationResponse is triggered when validation checks on the client's input fail (422).
// It sends the validation errors back to the client.
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

// editConflictResponse is sent when two processes attempt to edit the same record at the same time, causing a conflict (409).
// This helps prevent data corruption by informing the client to try again.
func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(w, r, http.StatusConflict, message)
}

// rateLimitExceededResponse is used when the client has exceeded the rate limit (429).
// This is a safeguard against overloading the system with requests.
func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

// invalidCredentialsResponse handles authentication failures, where the provided credentials are incorrect (401).
func (app *application) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

// invalidAuthenticationTokenResponse informs the client that the authentication token is invalid or missing (401).
// It sets the "WWW-Authenticate" header to prompt the client to provide valid credentials.
func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

// authenticationRequiredResponse is sent when the client attempts to access a resource that requires authentication but has not provided it (401).
func (app *application) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

// inactiveAccountResponse handles situations where the user's account is inactive and they attempt to access a restricted resource (403).
func (app *application) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

// notPermittedResponse is sent when a user tries to access a resource for which they lack the necessary permissions (403).
func (app *application) notPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}
