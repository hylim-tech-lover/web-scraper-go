package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type Quote struct {
	QuoteText string   `json:"quoteText"`
	Author    string   `json:"author"`
	AuthorUrl string   `json:"authorUrl"`
	Tag       []string `json:"tags"`
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
	c.OnHTML("div.quote", func(e *colly.HTMLElement) {

		// Use span.text as iteration condition to record every quote element.
		e.ForEach("span.text", func(i int, ne *colly.HTMLElement) {
			tags := make([]string, 0)

			// Remove special character from string
			trimmedString := regexp.MustCompile(`[^a-zA-Z0-9,.'-;!? ]+`).ReplaceAllString(ne.Text, "")

			// Same level as span.text
			// span --> small.author (small attribute with author as class name and child of span)
			author := e.ChildText("span small.author")
			authorUrl := fmt.Sprintf("https://quotes.toscrape.com%s", e.ChildAttr("span a", "href"))

			e.ForEach("div.tags a.tag", func(j int, tagText *colly.HTMLElement) {
				tags = append(tags, tagText.Text)
			})

			quote := Quote{
				QuoteText: strings.TrimSpace(trimmedString),
				Author:    author,
				AuthorUrl: authorUrl,
				Tag:       tags,
			}
			quotes = append(quotes, quote)
		})

	})

	// By default, web scrapper will extract quote from page 1 to 2.
	for pageNum := 1; pageNum <= 2; pageNum++ {
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
	err = os.WriteFile(fName, json, 0644)

	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
