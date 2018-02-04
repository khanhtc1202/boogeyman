package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/pkg/errors"
)

const DuckDuckGoBaseURL = "https://duckduckgo.com/?q="

type DuckDuckGoSpider struct {
	baseUrl string
	ofType  search_engine.SearchEngineType
}

func NewDuckDuckGoSpider() *DuckDuckGoSpider {
	return &DuckDuckGoSpider{
		baseUrl: DuckDuckGoBaseURL,
		ofType:  search_engine.DUCKDUCKGO,
	}
}

func (d *DuckDuckGoSpider) GetSearchEngineType() search_engine.SearchEngineType {
	return d.ofType
}

func (d *DuckDuckGoSpider) Query(keyword *domain.Keyword) (search_engine.Base, error) {

	doc := d.fetchFromInternet(keyword.String())
	resultsData := d.parseDocumentData(doc)
	if len(*resultsData) < 1 {
		return nil, errors.New("Error on query data from search engine!")
	}
	return search_engine.NewDuckDuckGo(keyword, resultsData), nil
}

func (d *DuckDuckGoSpider) fetchFromInternet(keyword string) *goquery.Document {
	doc, err := goquery.NewDocument(d.baseUrl + keyword)
	if err != nil {
		panic("Error fetching data from internet!")
	}
	return doc
}

func (d *DuckDuckGoSpider) parseDocumentData(doc *goquery.Document) *domain.ResultItems {
	resultsData := domain.EmptyResultItems()
	doc.Find(".result.results_links_deep.highlight_d").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("div .result__snippet").Text()
		time := "unknown"
		resultsData.Add(d.convertToDomain(title, url, description, time))
	})
	return resultsData
}

func (d *DuckDuckGoSpider) convertToDomain(
	title string,
	url string,
	description string,
	time string,
) *domain.ResultItem {
	return domain.NewResultItem(time, title, description, url)
}
