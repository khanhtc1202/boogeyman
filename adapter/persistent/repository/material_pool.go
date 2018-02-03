package repository

import (
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/pkg/errors"
)

type MaterialPool struct {
	collector        service.Collector
	searchEngineList *search_engine.SearchEngineList
	resultData       *search_engine.ResultDataPool
}

func NewMaterialPool(service service.Collector) *MaterialPool {
	resultData := search_engine.EmptyResultDataPool()
	return &MaterialPool{
		collector:        service,
		searchEngineList: nil,
		resultData:       resultData,
	}
}

func (m *MaterialPool) GetResultData() *search_engine.ResultDataPool {
	return m.resultData
}

func (m *MaterialPool) Fetch(keyword *domain.Keyword, searchEngineList *search_engine.SearchEngineList) {
	for _, searchEngineType := range *searchEngineList {
		resultData, err := m.collector.Query(searchEngineType, keyword)
		if err != nil {
			panic("Error on fetching data from search engine!")
		}
		m.resultData.Add(resultData)
	}
	m.searchEngineList = searchEngineList
}

func (m *MaterialPool) GetItemsFromSearchEngine(searchEngineType search_engine.SearchEngineType) (search_engine.Base, error) {
	if m.searchEngineList.Has(searchEngineType) {
		return m.resultData.FilterByEngineType(searchEngineType), nil
	} else {
		return nil, errors.New("Error on getting data from data pool, maybe not found type")
	}
}
