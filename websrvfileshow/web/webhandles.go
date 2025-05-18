package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"websrvfileshow/util"
)

// Handler that redirects all http requests to https
func redirectToHTTPS(responseWriter http.ResponseWriter, request *http.Request) {
	target := "https://" + request.Host + request.RequestURI
	http.Redirect(responseWriter, request, target, http.StatusMovedPermanently)
}

// Handler for processing an empty GET request, which returns HTML containing the content of the 'content.txt' file
func getMainHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	content, err := os.ReadFile(util.GetExecDir() + "content.txt")

	if err != nil {
		http.Error(responseWriter, "Failed to read file", http.StatusInternalServerError)
		return
	}

	regexpGroup := regexp.MustCompile(`([\r\n][ \t]*)([^ \t].+:)([ \t]*([\r\n]|$))`)
	content = []byte(regexpGroup.ReplaceAllString(string(content), `$1<span class="group">$2</span>$3`))

	regexpTodo := regexp.MustCompile(`([\r\n][ \t]*)((☐|-)[^\r\n]+)([\r\n]|$)`)
	content = []byte(regexpTodo.ReplaceAllString(string(content), `$1<span class="todo">$2</span>$4`))

	regexpDone := regexp.MustCompile(`([\r\n][ \t]*)((✔|\+)[^\r\n]+)([\r\n]|$)`)
	content = []byte(regexpDone.ReplaceAllString(string(content), `$1<span class="done">$2</span>$4`))

	regexpCancelled := regexp.MustCompile(`([\r\n][ \t]*)((✘|x)[^\r\n]+)([\r\n]|$)`)
	content = []byte(regexpCancelled.ReplaceAllString(string(content), `$1<span class="cancelled">$2</span>$4`))

	content = []byte(strings.ReplaceAll(string(content), "\r\n", "<br />"))

	htmlTemplate, err := os.ReadFile(util.GetExecDir() + "page_pattern.html")

	if err != nil {
		fmt.Println("error while reading file: ", err)
		htmlTemplate = []byte("error while reading file: " + err.Error())
	}

	html := strings.ReplaceAll(string(htmlTemplate), "${{content}}", string(content))

	responseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	responseWriter.Write([]byte(html))
}

/*
func getJsonHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"message": "Hello, JSON!",
		"status":  "success",
	}

	json.NewEncoder(responseWriter).Encode(data)
}
*/

// Handler for processing the post_file request, which writes the body to the file content.txt
func postFileHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// read body
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Failed to read body", http.StatusInternalServerError)
		return
	}
	defer request.Body.Close()

	err = os.WriteFile(util.GetExecDir()+"content.txt", body, 0644)
	if err != nil {
		http.Error(responseWriter, "Failed to write file", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Set("Content-Type", "application/json")

	data := map[string]string{
		"status": "success",
	}
	json.NewEncoder(responseWriter).Encode(data)
}
