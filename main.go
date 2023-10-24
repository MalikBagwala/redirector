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

type RedirectEntry struct {
	URL     string
	Tags    []string
	Title   string
	Display bool
	Logo    string
}

func redirectHandler(redirectMap map[string]RedirectEntry, cm *closestmatch.ClosestMatch, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=600")
		// Get the requested path from the URL
		requestedPath := r.URL.Path

		// If the requested path is "/", serve the HTML template
		if requestedPath == "/" {

			w.Header().Set("Content-Type", "text/html")
			data := struct {
				RedirectMap map[string]RedirectEntry
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
		http.Redirect(w, r, redirectMap[closestPath].URL, http.StatusFound)
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

	redirectMap := map[string]RedirectEntry{
		"/instagram":     {URL: os.Getenv("INSTAGRAM"), Title: "Instagram", Display: true},
		"/linkedin":      {URL: os.Getenv("LINKEDIN"), Title: "LinkedIn", Display: true},
		"/site":          {URL: os.Getenv("SITE"), Display: false},
		"/portfolio":     {URL: os.Getenv("SITE"), Title: "Portfolio", Display: true},
		"/landingpage":   {URL: os.Getenv("SITE"), Display: false},
		"/github":        {URL: os.Getenv("GITHUB"), Title: "GitHub", Display: true},
		"/gitlab":        {URL: os.Getenv("GITLAB"), Title: "GitLab", Display: true},
		"/resume":        {URL: os.Getenv("RESUME"), Title: "Resume", Display: true},
		"/cv":            {URL: os.Getenv("RESUME"), Display: false},
		"/stackoverflow": {URL: os.Getenv("STACKOVERFLOW"), Title: "StackOverflow", Display: true},
		"/tweet":         {URL: os.Getenv("TWEET"), Title: "Twitter", Display: true},
		"/threads":       {URL: os.Getenv("THREADS"), Title: "Threads", Display: true},
		"/email":         {URL: os.Getenv("MAIL"), Title: "Email", Display: true},
		"/company":       {URL: os.Getenv("COMPANY"), Title: "Company", Display: true},
		"/photo":         {URL: os.Getenv("PHOTO"), Title: "Avatar", Display: false},
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
