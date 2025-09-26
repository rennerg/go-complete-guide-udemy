package main

import "fmt"

type floatMap map[string]float64

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Twitter":  "https://www.twitter.com",
	}

	fmt.Println("Websites:", websites)

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println("Updated Websites:", websites)

	delete(websites, "Twitter")
	fmt.Println("After Deleting Twitter:", websites)

	if url, exists := websites["Google"]; exists {
		fmt.Println("Google URL:", url)
	} else {
		fmt.Println("Google not found")
	}

	for name, url := range websites {
		fmt.Printf("%s: %s\n", name, url)
	}

	courses := floatMap{
		"Go":         49.99,
		"Python":     39.99,
		"JavaScript": 29.99,
	}

	fmt.Println("Courses:", courses)
}
