package web

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Simple login and password validation
// Credentials should be stored in the gohelloworld_credentials environment variable,
// formatted as 'login1=password1;login2=password2'.
func checkCredentials(username, password string) bool {
	credentialsString := os.Getenv("websrvfileshow_credentials")

	if credentialsString == "" {
		fmt.Print("environment variable not found")
		return false
	}

	credentials := strings.Split(credentialsString, ";")

	for index := range credentials {
		loginPassword := strings.Split(credentials[index], "=")
		login := loginPassword[0]
		password1 := loginPassword[1]

		if username == login && password == password1 {
			return true
		}
	}

	return false
}

func basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Unauthorized")
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[len("Basic "):])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid authorization header")
			return
		}

		parts := strings.SplitN(string(payload), ":", 2)
		if len(parts) != 2 || !checkCredentials(parts[0], parts[1]) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
