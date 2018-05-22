package repository

import (
	"fmt"

	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/pkg/errors"
)

type MaterialPool struct {
	collectors       []service.Collector
	searchEngineList *search_engine.SearchEngineList
	resultData       *search_engine.ResultDataPool
}

func NewMaterialPool(services []service.Collector) *MaterialPool {
	searchEngineList := search_engine.EmptySearchEngineList()
	for _, searchEngine := range services {
		searchEngineList.Add(searchEngine.GetSearchEngineType())
	}
	resultData := search_engine.EmptyResultDataPool()
	return &MaterialPool{
		collectors:       services,
		searchEngineList: searchEngineList,
		resultData:       resultData,
	}
}

func (m *MaterialPool) GetResultData() *search_engine.ResultDataPool {
	return m.resultData
}

func (m *MaterialPool) GetSearchEngineList() *search_engine.SearchEngineList {
	return m.searchEngineList
}

func (m *MaterialPool) Fetch(keyword *domain.Keyword) error {
	for _, collector := range m.collectors {
		resultData, err := collector.Query(keyword)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Error on fetching data from search engine! \n"))
		}
		m.resultData.Add(resultData)
	}
	return nil
}

func (m *MaterialPool) GetItemsFromSearchEngine(searchEngineType search_engine.SearchEngineType) (search_engine.SearchEngine, error) {
	if m.searchEngineList.Has(searchEngineType) {
		return m.resultData.FilterByEngineType(searchEngineType), nil
	} else {
		return nil, errors.New("Error on getting data from data pool, maybe not found type")
	}
}
