package service

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/pkg/errors"
)

const AskBaseURL = "https://www.ask.com/web?q="

type AskSpider struct {
	baseUrl string
	ofType  domain.SearchEngineType
}

func NewAskSpider() *AskSpider {
	return &AskSpider{
		baseUrl: AskBaseURL,
		ofType:  domain.ASK,
	}
}

func (a *AskSpider) GetSearchEngineType() domain.SearchEngineType {
	return a.ofType
}

func (a *AskSpider) Query(keyword *domain.Keyword) (*domain.SearchEngine, error) {

	doc := a.fetchFromInternet(keyword.String())
	resultsData := a.parseDocumentData(doc)
	if len(*resultsData) < 1 {
		return nil, errors.New("Error on query data from search engine (Ask)!")
	}
	return domain.NewSearchEngine(a.ofType, resultsData), nil
}

func (a *AskSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(a.baseUrl + keyword)
	if err != nil {
		fmt.Println("Error fetching data from ask.com!")
		os.Exit(1)
	}
	return doc
}

func (a *AskSpider) parseDocumentData(doc *goquery.Document) *domain.QueryResult {
	resultsData := domain.EmptyQueryResult()
	doc.Find(".PartialSearchResults-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("p.PartialSearchResults-item-abstract").Text()
		time := "unknown"
		resultsData.Add(a.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (a *AskSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.ResultItem {
	return domain.NewResultItem(time, title, description, url)
}
