package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/pkg/errors"
)

const AskBaseURL = "https://www.ask.com/web?q="

type AskSpider struct {
	baseUrl string
	ofType  search_engine.SearchEngineType
}

func NewAskSpider() *AskSpider {
	return &AskSpider{
		baseUrl: AskBaseURL,
		ofType:  search_engine.ASK,
	}
}

func (b *AskSpider) GetSearchEngineType() search_engine.SearchEngineType {
	return b.ofType
}

func (b *AskSpider) Query(keyword *domain.Keyword) (search_engine.Base, error) {

	doc := b.fetchFromInternet(keyword.String())
	resultsData := b.parseDocumentData(doc)
	if len(*resultsData) < 1 {
		return nil, errors.New("Error on query data from search engine!")
	}
	return search_engine.NewAsk(keyword, resultsData), nil
}

func (b *AskSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(b.baseUrl + keyword)
	if err != nil {
		panic("Error fetching data from internet!")
	}
	return doc
}

func (b *AskSpider) parseDocumentData(doc *goquery.Document) *domain.ResultItems {
	resultsData := domain.EmptyResultItems()
	doc.Find(".PartialSearchResults-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("p.PartialSearchResults-item-abstract").Text()
		time := "unknown"
		resultsData.Add(b.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (b *AskSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.ResultItem {
	return domain.NewResultItem(time, title, description, url)
}
