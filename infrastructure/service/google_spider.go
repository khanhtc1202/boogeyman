package service

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/pkg/errors"
)

const GoogleBaseURL = "https://www.google.com/search?q="

type GoogleSpider struct {
	baseUrl string
	ofType  search_engine.SearchEngineType
}

func NewGoogleSpider() *GoogleSpider {
	return &GoogleSpider{
		baseUrl: GoogleBaseURL,
		ofType:  search_engine.GOOGLE,
	}
}

func (g *GoogleSpider) GetSearchEngineType() search_engine.SearchEngineType {
	return g.ofType
}

func (g *GoogleSpider) Query(keyword *domain.Keyword) (search_engine.Base, error) {

	doc := g.fetchFromInternet(keyword.String())
	resultsData := g.parseDocumentData(doc)
	if len(*resultsData) < 1 {
		return nil, errors.New("Error on query data from search engine!")
	}
	return search_engine.NewGoogle(keyword, resultsData), nil
}

func (g *GoogleSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(g.baseUrl + keyword)
	if err != nil {
		panic("Error fetching data from internet!")
	}
	return doc
}

func (g *GoogleSpider) parseDocumentData(doc *goquery.Document) *domain.ResultItems {
	resultsData := domain.EmptyResultItems()
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
