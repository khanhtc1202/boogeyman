package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/pkg/errors"
)

const BingBaseURL = "https://www.bing.com/search?q="

type BingSpider struct {
	baseUrl string
}

func NewBingSpider() *BingSpider {
	return &BingSpider{
		baseUrl: BingBaseURL,
	}
}

func (b *BingSpider) Query(searchEngineType search_engine.SearchEngineType, keyword *domain.Keyword) (search_engine.Base, error) {

	doc := b.fetchFromInternet(keyword.String())
	resultsData := b.parseDocumentData(doc)
	if len(*resultsData) < 1 {
		return nil, errors.New("Error on query data from search engine!")
	}
	return search_engine.NewBing(keyword, resultsData), nil
}

func (b *BingSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(b.baseUrl + keyword)
	if err != nil {
		panic("Error fetching data from internet!")
	}
	return doc
}

func (b *BingSpider) parseDocumentData(doc *goquery.Document) *domain.ResultItems {
	resultsData := domain.EmptyResultItems()
	doc.Find(".b_algo").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("p").Text()
		time := "unknown"
		resultsData.Add(b.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (b *BingSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.ResultItem {
	return domain.NewResultItem(time, title, description, url)
}
