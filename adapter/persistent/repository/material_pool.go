package repository

import (
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/pkg/errors"
)

type MaterialPool struct {
	collectors       []service.Collector
	searchEngineList *domain.SearchEngineList
	resultData       *domain.QueryResultPool
}

func NewMaterialPool(services []service.Collector) *MaterialPool {
	searchEngineList := domain.EmptySearchEngineList()
	for _, searchEngine := range services {
		searchEngineList.Add(searchEngine.GetSearchEngineType())
	}
	resultData := domain.EmptyQueryResultPool()
	return &MaterialPool{
		collectors:       services,
		searchEngineList: searchEngineList,
		resultData:       resultData,
	}
}

func (m *MaterialPool) GetResultData() *domain.QueryResultPool {
	return m.resultData
}

func (m *MaterialPool) GetSearchEngineList() *domain.SearchEngineList {
	return m.searchEngineList
}

func (m *MaterialPool) Fetch(keyword *domain.Keyword) error {
	for _, collector := range m.collectors {
		resultData, err := collector.Query(keyword)
		if err != nil {
			return errors.Wrap(err, "Error on fetching data from search engine! \n")
		}
		m.resultData.Add(resultData)
	}
	return nil
}

func (m *MaterialPool) GetItemsFromSearchEngine(
	searchEngineType domain.SearchEngineType,
) (*domain.SearchEngine, error) {
	if m.searchEngineList.Has(searchEngineType) {
		return m.resultData.FilterByEngineType(searchEngineType), nil
	} else {
		return nil, errors.New("Error on getting data from data pool, maybe not found type")
	}
}
