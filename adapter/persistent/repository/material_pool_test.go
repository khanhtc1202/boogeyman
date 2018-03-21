package repository_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
)

type CollectorMock struct {
	service.Collector
}

func (c *CollectorMock) GetSearchEngineType() search_engine.SearchEngineType {
	return search_engine.GOOGLE
}

func (c *CollectorMock) Query(keyword *domain.Keyword) (search_engine.SearchEngine, error) {
	return search_engine.NewGoogle(keyword, fakeResultList()), nil
}

func TestMaterialPool_Fetch(t *testing.T) {
	keyword := domain.NewKeyword("sample")

	materialPool := repository.NewMaterialPool([]service.Collector{&CollectorMock{}})

	materialPool.Fetch(keyword)
	if len(*materialPool.GetResultData()) != len(*materialPool.GetSearchEngineList()) {
		t.Fatal("Fail on test fetch data from search engine")
	}
}

func TestMaterialPool_GetItemsBySearchEngine(t *testing.T) {
	keyword := domain.NewKeyword("sample")

	materialPool := repository.NewMaterialPool([]service.Collector{&CollectorMock{}})
	materialPool.Fetch(keyword)

	searchResult, err := materialPool.GetItemsFromSearchEngine(search_engine.GOOGLE)
	if err != nil {
		t.Fatal("Fail on test get items from search engine")
	}
	if searchResult == nil || searchResult.Type() != search_engine.GOOGLE {
		t.Fatal("Fail type of return value on query from search engine")
	}
	if len(*searchResult.GetResults()) != 2 {
		t.Fatal("Fail on query from search engine, maybe fail by network connection")
	}
}

func fakeResultList() *domain.ResultItems {
	fakeResult1 := domain.NewResultItem("timestamp", "title 1", "my desc", "http://sample.com/acv")
	fakeResult2 := domain.NewResultItem("timestamp", "title 2", "my desc", "http://sample.com/123")

	results := domain.EmptyResultItems()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}
