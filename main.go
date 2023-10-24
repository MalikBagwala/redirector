package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/schollz/closestmatch"
	"golang.org/x/exp/maps"
)

func redirectHandler(redirectMap map[string]string, cm *closestmatch.ClosestMatch, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=600")
		// Get the requested path from the URL
		requestedPath := r.URL.Path

		// If the requested path is "/", serve the HTML template
		if requestedPath == "/" {

			w.Header().Set("Content-Type", "text/html")
			data := struct {
				RedirectMap map[string]string
			}{
				RedirectMap: redirectMap,
			}
			err := tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Find the closest path and perform redirection
		closestPath := cm.Closest(requestedPath)
		http.Redirect(w, r, redirectMap[closestPath], http.StatusFound)
	}
}

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}
	// Parse the HTML template
	tmpl, teplErr := template.ParseFiles("template.html")
	if teplErr != nil {
		log.Fatal(teplErr)
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
	http.HandleFunc("/", redirectHandler(redirectMap, cm, tmpl))

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
