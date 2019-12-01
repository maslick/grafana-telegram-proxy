package main

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

func getPort() string {
	var port = getEnv("PORT", "8080")
	return ":" + port
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type handler func(w http.ResponseWriter, r *http.Request)

func basicAuth(pass handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validateUsernamePassword(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		pass(w, r)
	}
}

func validateUsernamePassword(username, password string) bool {
	if username == getEnv("USERNAME", "") && password == getEnv("PASSWORD", "") {
		return true
	}
	return false
}
