package service

import (
	"strings"

	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/pkg/errors"
)

const GoogleBaseURL = "https://www.google.com/search?q="

type GoogleSpider struct {
	baseUrl string
	ofType  domain.SearchEngineType
}

func NewGoogleSpider() *GoogleSpider {
	return &GoogleSpider{
		baseUrl: GoogleBaseURL,
		ofType:  domain.GOOGLE,
	}
}

func (g *GoogleSpider) GetSearchEngineType() domain.SearchEngineType {
	return g.ofType
}

func (g *GoogleSpider) Query(keyword *domain.Keyword) (*domain.SearchEngine, error) {

	doc := g.fetchFromInternet(keyword.String())
	resultsData := g.parseDocumentData(doc)
	if len(*resultsData) < 1 {
		return nil, errors.New("Error on query data from search engine (Google)!")
	}
	return domain.NewSearchEngine(g.ofType, resultsData), nil
}

func (g *GoogleSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(g.baseUrl + keyword)
	if err != nil {
		fmt.Println("Error fetching data from google.com!")
		os.Exit(1)
	}
	return doc
}

func (g *GoogleSpider) parseDocumentData(doc *goquery.Document) *domain.QueryResult {
	resultsData := domain.EmptyQueryResult()
	doc.Find(".g").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find(".st").Text()
		time := "unknown"
		resultsData.Add(g.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (g *GoogleSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.ResultItem {
	url = strings.Replace(url, "/url?q=", "", -1)
	return domain.NewResultItem(time, title, description, url)
}
