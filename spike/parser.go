package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("https://www.bing.com/search?q=something")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".b_algo").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("p").Text()
		fmt.Printf("=>> Result %d: \n %s - %s \n %s\n", i, title, url, description)
	})
}

func main() {
	ExampleScrape()
}
