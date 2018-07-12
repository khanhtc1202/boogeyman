package service

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/PuerkitoBio/goquery"
	"os"
	"fmt"
)

const YahooBaseURL = "https://search.yahoo.com/search?p="

type YahooSpider struct {
	baseUrl string
	ofType	domain.SearchEngineType
}

func NewYahooSpider() *YahooSpider {
	return &YahooSpider{
		baseUrl: YahooBaseURL,
		ofType:	domain.YAHOO,
	}
}

func (y *YahooSpider) GetSearchEngineType() domain.SearchEngineType {
	return y.ofType
}

func (y *YahooSpider) Query(keyword *domain.Keyword) (*domain.SearchEngine, error) {
	doc := y.fetchFromInternet(keyword.String())
	resultsData := y.parseDocumentData(doc)
	return domain.NewSearchEngine(y.ofType, resultsData), nil
}

func (y *YahooSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(y.baseUrl + keyword)
	if err != nil {
		fmt.Println("Error fetching data from google.com!")
		os.Exit(1)
	}
	return doc
}

func (y *YahooSpider) parseDocumentData(doc *goquery.Document) *domain.QueryResult {
	resultsData := domain.EmptyQueryResult()
	doc.Find(".algo-sr").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("p").Text()
		time := "unknown"
		resultsData.Add(y.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (y *YahooSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.ResultItem {
	return domain.NewResultItem(time, title, description, url)
}
