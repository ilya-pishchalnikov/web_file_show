package web

import (
	"fmt"
	"net/http"
)

// Initializes and starts the HTTP server
func StartServer(port string, cert string, certKey string) {

	go func() {
		http.ListenAndServe(":80", http.HandlerFunc(redirectToHTTPS))
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", getMainHandler)
	//mux.HandleFunc("/get_json", getJsonHandler)
	mux.HandleFunc("/post_file", postFileHandler)

	fmt.Println("Server listening on port", port)

	err := http.ListenAndServeTLS(port, cert, certKey, basicAuth(mux))
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
}
