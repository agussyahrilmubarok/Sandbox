package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"greenlight.id/internal/validator"
)

// readIDParam retrieves the "id" parameter from the request context and converts it to an int64.
// It returns an error if the ID is invalid (non-numeric or less than 1).
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

// envelope is a custom type used for structuring JSON responses.
type envelope map[string]interface{}

// writeJSON formats the provided data as JSON and writes it to the response writer.
// It sets the appropriate content type and status code. If an error occurs during marshaling, it returns the error.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t") // Marshal data into indented JSON
	if err != nil {
		return err
	}

	js = append(js, '\n') // Append a newline for better readability

	// Set headers from the provided map
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	w.WriteHeader(status)                              // Set the HTTP status code
	w.Write(js)                                        // Write the JSON response

	return nil
}

// readJSON decodes JSON from the request body into the provided destination struct.
// It handles various JSON decoding errors and returns a meaningful error message if issues arise.
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576                                    // Limit the size of the request body to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes)) // Enforce the size limit
	dec := json.NewDecoder(r.Body)                           // Create a new JSON decoder
	dec.DisallowUnknownFields()                              // Prevent decoding of unknown fields

	// Attempt to decode the JSON body into the destination
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		// Handle specific types of decoding errors
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			panic(err) // Panic if there is an invalid unmarshal error

		default:
			return err // Return the original error if it's unrecognized
		}
	}

	// Ensure that the body contains only a single JSON value
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}

// readString retrieves a string value from the query string parameters.
// If the parameter is not present, it returns a default value.
func (app *application) readString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

// readCSV retrieves a comma-separated string from the query string parameters and splits it into a slice.
// If the parameter is absent, it returns a default slice.
func (app *application) readCSV(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}

// readInt retrieves an integer value from the query string parameters.
// It returns a default value if the parameter is missing or invalid, and records validation errors.
func (app *application) readInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value") // Add validation error if conversion fails
		return defaultValue
	}

	return i
}

// background runs a function in a separate goroutine.
// It handles any panic that occurs during execution and logs the error.
func (app *application) background(fn func()) {
	app.wg.Add(1) // Increment the wait group counter

	go func() {
		defer app.wg.Done() // Decrement the counter when the goroutine completes

		defer func() {
			if err := recover(); err != nil {
				app.logger.PrintError(fmt.Errorf("%s", err), nil) // Log any panic that occurs
			}
		}()

		fn() // Execute the function
	}()
}
