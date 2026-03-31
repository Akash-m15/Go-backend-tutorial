package auth

import (
	"errors"
	"net/http"
	"strings"
)

//Example:  Authorization: ApiKey {api_key}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("No auth token found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("Malformed Auth Header")
	}

	return vals[1], nil
}
