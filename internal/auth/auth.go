package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const APIKEY_FORMAT = "ApiKey"

// GetAPIKey extracts an API key from the headers of an HTTP request
// Example:
// Authorization: ApiKey {actual apikey}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Missing header Authorization")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New(fmt.Sprintf("Malformed authorization header. Expected format: %v %v", APIKEY_FORMAT, "example_key"))
	}

	if vals[0] != APIKEY_FORMAT {
		return "", errors.New(fmt.Sprintf("Malformed first part of authorization header. Expected format: %v %v", APIKEY_FORMAT, "example_key"))
	}

	return vals[1], nil
}
