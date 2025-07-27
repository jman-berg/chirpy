package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	api_key_string := headers.Get("Authorization")

	if api_key_string == "" {
		return "", errors.New("No token found in authorization header")
	}

	return strings.TrimPrefix(api_key_string, "ApiKey "), nil

}
