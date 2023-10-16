package main

import (
	"os"
)

var redirectMap = map[string]string{
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
