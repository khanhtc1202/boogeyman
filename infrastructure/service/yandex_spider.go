package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/pkg/errors"
)

const YandexBaseURL = "https://yandex.com/search/?text="

type YandexSpider struct {
	baseUrl string
	ofType  search_engine.SearchEngineType
}

func NewYandexSpider() *YandexSpider {
	return &YandexSpider{
		baseUrl: YandexBaseURL,
		ofType:  search_engine.YANDEX,
	}
}

func (b *YandexSpider) GetSearchEngineType() search_engine.SearchEngineType {
	return b.ofType
}

func (b *YandexSpider) Query(keyword *domain.Keyword) (search_engine.SearchEngine, error) {

	doc := b.fetchFromInternet(keyword.String())
	resultsData := b.parseDocumentData(doc)
	if len(*resultsData) < 1 {
		return nil, errors.New("Error on query data from search engine!")
	}
	return search_engine.NewYandex(keyword, resultsData), nil
}

func (b *YandexSpider) fetchFromInternet(keyword string) *goquery.Document {
	// TODO yandex search engine catch request from boogeyman as an automate req
	doc, err := goquery.NewDocument(b.baseUrl + keyword)
	if err != nil {
		panic("Error fetching data from internet!")
	}
	return doc
}

func (b *YandexSpider) parseDocumentData(doc *goquery.Document) *domain.ResultItems {
	resultsData := domain.EmptyResultItems()
	doc.Find(".serp-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".organic__url-text").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("div.text-container").Text()
		time := "unknown"
		resultsData.Add(b.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (b *YandexSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.ResultItem {
	return domain.NewResultItem(time, title, description, url)
}
