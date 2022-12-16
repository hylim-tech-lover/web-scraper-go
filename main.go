package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gocolly/colly"
)

type Quote struct {
	QuoteText string `json:"quoteText"`
}

func main() {
	quotes := make([]Quote, 0)

	// Instantiate default collector
	c := colly.NewCollector(
		// Allow requests to
		colly.AllowedDomains("quotes.toscrape.com", "www.quotes.toscrape.com"),
	)

	// Extract Quote file
	// TODO: Try to get author as well
	c.OnHTML(".quote span.text", func(e *colly.HTMLElement) {

		// Remove special character from string
		trimmedString := regexp.MustCompile(`[^a-zA-Z0-9,. ]+`).ReplaceAllString(e.Text, "")

		quote := Quote{
			QuoteText: trimmedString,
		}
		quotes = append(quotes, quote)
	})

	for i := 0; i <= 1; i++ {
		pageNum := i + 1
		url := fmt.Sprintf("https://quotes.toscrape.com/page/%d", pageNum)
		c.Visit(url)
	}

	// Debugging purpose
	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "  ")
	// enc.Encode(quotes)

	writeJson(quotes)
}

func createFolder(path string) {
	folderPath := filepath.Join(".", path)
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create folder. Program will exit")
		return
	}
}

func writeJson(data []Quote) {

	const folderName = "generated"
	createFolder(folderName)

	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}

	fName := fmt.Sprintf("%s/quote.json", folderName)
	// Create file if not yet exist
	err = ioutil.WriteFile(fName, json, 0644)

	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
