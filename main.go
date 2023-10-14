package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=600")
	redirectMap := map[string]string{
		"/insta":   "https://www.instagram.com/straightfllush/",
		"/lkdin":   "https://www.linkedin.com/in/malikbagwala/",
		"/site":    "https://maalik.dev",
		"/github":  "https://github.com/MalikBagwala",
		"/gitlab":  "http://gitlab.com/MalikBagwala",
		"/resume":  "https://drive.google.com/file/d/1vBwDKM0bFcRXGe4jkbnhIDRQU_cJZD9j/view",
		"/stack":   "https://stackoverflow.com/users/10177043/malik-bagwala",
		"/tweet":   "https://twitter.com/MalikBagwala",
		"/threads": "https://www.threads.net/@straightfllush",
	}

	// Get the requested path from the URL
	requestedPath := r.URL.Path

	// Check if the requested path exists in the redirect map
	if redirectURL, ok := redirectMap[requestedPath]; ok {
		// Redirect to the specified URL with a 302 status code (Temporary Redirect)
		http.Redirect(w, r, redirectURL, http.StatusFound)
		return
	}

	// If the requested path is "/", return the redirect map as JSON
	if requestedPath == "/" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(redirectMap)
		return
	}

	// If the requested path is not found, return a 404 Not Found response
	http.NotFound(w, r)
}

func main() {
	// Create a new HTTP server and set up the redirectHandler for all routes
	http.HandleFunc("/", redirectHandler)

	// Specify the port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start the server
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}