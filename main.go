package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/schollz/closestmatch"
	"golang.org/x/exp/maps"
)

func redirectHandler(redirectMap map[string]string, cm *closestmatch.ClosestMatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=600")
		// Get the requested path from the URL
		requestedPath := r.URL.Path
		// If the requested path is "/", return the redirect map as JSON
		if requestedPath == "/" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(redirectMap)
			return
		}
		// Find the closest path
		// eg. /insta -> /instagram
		// eg. /intagram -> /instagram
		closestPath := cm.Closest(requestedPath)
		http.Redirect(w, r, redirectMap[closestPath], http.StatusFound)
	}
}
func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}
	redirectMap := map[string]string{
		"/instagram":     os.Getenv("INSTAGRAM"),
		"/linkedin":      os.Getenv("LINKEDIN"),
		"/site":          os.Getenv("SITE"),
		"/portfolio":     os.Getenv("SITE"),
		"/landingpage":   os.Getenv("SITE"),
		"/github":        os.Getenv("GITHUB"),
		"/gitlab":        os.Getenv("GITLAB"),
		"/resume":        os.Getenv("RESUME"),
		"/cv":            os.Getenv("RESUME"),
		"/stackoverflow": os.Getenv("STACKOVERFLOW"),
		"/tweet":         os.Getenv("TWEET"),
		"/threads":       os.Getenv("THREADS"),
		"/email":         os.Getenv("MAIL"),
		"/company":       os.Getenv("COMPANY"),
		"/photo":         os.Getenv("PHOTO"),
	}

	// Convert redirect map to list of its keys
	wordsToTest := maps.Keys(redirectMap)
	cm := closestmatch.New(wordsToTest, []int{1})

	// Create a new HTTP server and set up the redirectHandler for all routes
	http.HandleFunc("/", redirectHandler(redirectMap, cm))

	// Specify the port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Start the server
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
