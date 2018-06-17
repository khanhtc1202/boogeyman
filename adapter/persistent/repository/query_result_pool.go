package repository

import (
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/pkg/errors"
)

type QueryResultPool struct {
	collectors       []service.Collector
	searchEngineList *domain.SearchEngineList
}

func NewResultPool(services []service.Collector) *QueryResultPool {
	searchEngineList := domain.EmptySearchEngineList()
	for _, searchEngine := range services {
		searchEngineList.Add(searchEngine.GetSearchEngineType())
	}
	return &QueryResultPool{
		collectors:       services,
		searchEngineList: searchEngineList,
	}
}

func (m *QueryResultPool) GetSearchEngineList() *domain.SearchEngineList {
	return m.searchEngineList
}

func (m *QueryResultPool) FetchData(
	keyword *domain.Keyword,
) (*domain.QueryResultPool, error) {
	resultPool := domain.EmptyQueryResultPool()
	for _, collector := range m.collectors {
		resultData, err := collector.Query(keyword)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetching data from search engine! \n")
		}
		resultPool.Add(resultData)
	}
	return resultPool, nil
}
