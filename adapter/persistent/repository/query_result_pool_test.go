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

func TestMaterialPool_FetchData(t *testing.T) {
	keyword := domain.NewKeyword("sample")

	resultPoolRepo := repository.NewResultPool([]service.Collector{&CollectorMock{}})

	resultPool, _ := resultPoolRepo.FetchData(keyword)
	if len(*resultPool) != 1 {
		t.Fatal("Fail on test fetch data from search engine")
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
