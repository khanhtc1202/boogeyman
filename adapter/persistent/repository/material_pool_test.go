package repository_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
)

type CollectorMock struct {
	service.Collector
}

func (c *CollectorMock) GetSearchEngineType() domain.SearchEngineType {
	return domain.GOOGLE
}

func (c *CollectorMock) Query(keyword *domain.Keyword) (*domain.SearchEngine, error) {
	return domain.NewSearchEngine(domain.GOOGLE, fakeResultList()), nil
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

	searchResult, err := materialPool.GetItemsFromSearchEngine(domain.GOOGLE)
	if err != nil {
		t.Fatal("Fail on test get items from search engine")
	}
	if searchResult == nil || searchResult.Type() != domain.GOOGLE {
		t.Fatal("Fail type of return value on query from search engine")
	}
	if len(*searchResult.GetQueryResults()) != 2 {
		t.Fatal("Fail on query from search engine, maybe fail by network connection")
	}
}

func fakeResultList() *domain.QueryResult {
	fakeResult1 := domain.NewResultItem("timestamp", "title 1", "my desc", "http://sample.com/acv")
	fakeResult2 := domain.NewResultItem("timestamp", "title 2", "my desc", "http://sample.com/123")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}
