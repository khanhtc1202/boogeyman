package main

import (
	"fmt"
	"log"

	"errors"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func ExampleScrape(queryString string) {
	doc, err := goquery.NewDocument("https://www.bing.com/search?q=" + queryString)
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

func ParseCommandParams() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("Missing command params")
	}
	return os.Args[1:], nil
}

func main() {
	cmdParams, err := ParseCommandParams()
	if err != nil {
		fmt.Println(err)
	}
	ExampleScrape(cmdParams[0])
}
