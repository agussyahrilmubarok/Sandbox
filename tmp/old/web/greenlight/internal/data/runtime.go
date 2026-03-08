package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// `ErrInvalidRuntimeFormat` is an error that occurs when the runtime format is invalid during unmarshaling.
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// `Runtime` is a custom type based on `int32`, representing the duration of a movie in minutes.
type Runtime int32

// `MarshalJSON` implements the `json.Marshaler` interface, converting a `Runtime` value into a JSON string.
// The runtime is formatted as "{value} mins" (e.g., "120 mins").
func (r Runtime) MarshalJSON() ([]byte, error) {
	// Convert the runtime to a string in the format "X mins".
	jsonValue := fmt.Sprintf("%d mins", r)

	// Add double quotes around the string to properly format it as JSON.
	quotedJSONValue := strconv.Quote(jsonValue)

	// Return the JSON-encoded byte slice.
	return []byte(quotedJSONValue), nil
}

// `UnmarshalJSON` implements the `json.Unmarshaler` interface, converting a JSON string into a `Runtime` value.
// The expected format is "{value} mins" (e.g., "120 mins").
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// Remove the surrounding quotes from the JSON string.
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		// If the value can't be unquoted, return an error.
		return ErrInvalidRuntimeFormat
	}

	// Split the string into two parts: the number and the "mins" suffix.
	parts := strings.Split(unquotedJSONValue, " ")
	// Check if the split is exactly two parts and the second part is "mins".
	if len(parts) != 2 || parts[1] != "mins" {
		// If the format is incorrect, return an error.
		return ErrInvalidRuntimeFormat
	}

	// Parse the first part (the number of minutes) into an integer.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		// If parsing fails, return an error indicating an invalid format.
		return ErrInvalidRuntimeFormat
	}

	// Assign the parsed value to the `Runtime` instance.
	*r = Runtime(i)
	return nil
}
