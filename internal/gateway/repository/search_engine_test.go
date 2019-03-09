package repository

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/internal/domain"
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
	enginesRepo := SearchEngines(collectors)

	resultPool, _ := enginesRepo.FetchData(keyword)
	if len(*resultPool) != 1 {
		t.Fatal("Fail on test fetch data from search engine")
	}
}

func TestSearchEngines_AddEnginesByType_AddAllEngineByDefault(t *testing.T) {
	collectors := service.EmptyCollectorList()
	enginesRepo := SearchEngines(collectors)

	err := enginesRepo.AddEnginesByType(domain.UNKNOWN_ENGINE)
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(*enginesRepo.collectors) != 4 {
		t.Fatal("Error not all engines been added")
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
