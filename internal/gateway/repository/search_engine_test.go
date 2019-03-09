package repository_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/gateway/repository"
	"github.com/khanhtc1202/boogeyman/internal/gateway/service"
)

type CollectorMock struct {
	service.Collector
}

func (c *CollectorMock) GetSearchEngineType() domain.SearchEngineType {
	return domain.GOOGLE
}

func (c *CollectorMock) Query(keyword domain.Keyword) (*domain.SearchEngine, error) {
	return domain.NewSearchEngine(domain.GOOGLE, fakeResultList()), nil
}

func TestMaterialPool_FetchData(t *testing.T) {
	keyword := domain.NewKeyword("sample")

	collectors := service.EmptyCollectorList()
	collectors.Add(&CollectorMock{})
	resultPoolRepo := repository.SearchEngines(collectors)

	resultPool, _ := resultPoolRepo.FetchData(keyword)
	if len(*resultPool) != 1 {
		t.Fatal("Fail on test fetch data from search engine")
	}
}

func fakeResultList() *domain.QueryResults {
	fakeResult1 := domain.NewResultItem("timestamp", "title 1", "my desc", "http://sample.com/acv")
	fakeResult2 := domain.NewResultItem("timestamp", "title 2", "my desc", "http://sample.com/123")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}