package validator

import "regexp"

// `EmailRX` is a regular expression used to validate the format of an email address.
// It follows the standard email format as per RFC 5322 specification.
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// `Validator` is a struct that holds a map of validation errors.
// Each key in the map corresponds to a field or category, and the value is the associated error message.
type Validator struct {
	Errors map[string]string
}

// `New` is a helper function that creates a new `Validator` instance with an empty errors map.
// This is useful for starting a fresh validation process.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// `Valid` returns true if the errors map is empty, indicating no validation errors have been recorded.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// `AddError` adds an error message to the errors map, but only if an entry for the given key doesn't already exist.
// This ensures that no duplicate error messages are added for the same field.
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// `Check` adds an error message to the errors map only if the validation check (represented by `ok`) fails (i.e., is `false`).
// This simplifies condition-based validation, where an error message is added if the condition fails.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// `In` returns true if a given string value is present in a list of strings.
// This is used to validate if a value belongs to a predefined set of allowed values.
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// `Matches` returns true if a given string matches a regular expression pattern.
// This is useful for validating specific formats like email addresses, phone numbers, etc.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// `Unique` returns true if all the string values in a slice are unique.
// This checks if there are any duplicate values in the slice.
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}
