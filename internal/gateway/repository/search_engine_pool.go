package repository

import (
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/gateway/service"
	"github.com/pkg/errors"
)

type SearchEnginePool struct {
	collectors []service.Collector
}

func NewSearchEnginePool(
	services []service.Collector,
) *SearchEnginePool {
	return &SearchEnginePool{
		collectors: services,
	}
}

func (m *SearchEnginePool) FetchData(
	keyword domain.Keyword,
) (*domain.SearchEnginePool, error) {
	searchEnginePool := domain.EmptySearchEnginePool()
	resultsChan := make(chan *domain.SearchEngine, len(m.collectors))
	errChan := make(chan error)

	for _, collector := range m.collectors {
		go func(collector service.Collector) {
			resultData, err := collector.Query(keyword)
			if err != nil {
				errChan <- err
			}
			resultsChan <- resultData
		}(collector)
	}

	for {
		select {
		case err := <-errChan:
			return nil, errors.Wrap(err, "Error on fetching data from search engine! \n")
		case resultData := <-resultsChan:
			searchEnginePool.Add(resultData)
			if len(*searchEnginePool) == len(m.collectors) {
				return searchEnginePool, nil
			}
		}
	}
}
