package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey: extracts APIKey from the headers of a http requests
// Example of a header it looks out for
// Authorization: ApiKey {insert the apikey here}
func GetAPIKey(headers http.Header) (string, error) {
  val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	
	return vals[1], nil
}