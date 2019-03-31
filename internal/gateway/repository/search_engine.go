package repository

import (
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/gateway/service"
	collectorPool "github.com/khanhtc1202/boogeyman/internal/infrastructure/service"
	"github.com/pkg/errors"
)

type searchEngines struct {
	collectors *service.CollectorList
}

func SearchEngines(
	services *service.CollectorList,
) *searchEngines {
	return &searchEngines{
		collectors: services,
	}
}

func (s *searchEngines) FetchData(
	keyword domain.Keyword,
) (*domain.SearchEnginePool, error) {
	searchEnginePool := domain.EmptySearchEnginePool()
	resultsChan := make(chan *domain.SearchEngine, len(*s.collectors))
	errChan := make(chan error)

	for _, collector := range *s.collectors {
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
			if len(*searchEnginePool) == len(*s.collectors) {
				return searchEnginePool, nil
			}
		}
	}
}

func (s *searchEngines) AddEnginesByType(engineType domain.SearchEngineType) error {
	if s.has(engineType) {
		return errors.New("Search engine already added!\n")
	}

	switch engineType {
	case domain.ASK:
		s.collectors.Add(collectorPool.NewAskSpider())
	case domain.BING:
		s.collectors.Add(collectorPool.NewBingSpider())
	case domain.YAHOO:
		s.collectors.Add(collectorPool.NewYahooSpider())
	case domain.GOOGLE:
		s.collectors.Add(collectorPool.NewGoogleSpider())
	default:
		// return collect from all search engine by default
		s.addAllEngines()
	}

	return nil
}

func (s *searchEngines) addAllEngines() {
	collector := service.EmptyCollectorList()
	collector.Add(collectorPool.NewAskSpider())
	collector.Add(collectorPool.NewBingSpider())
	collector.Add(collectorPool.NewYahooSpider())
	collector.Add(collectorPool.NewGoogleSpider())

	s.collectors = collector
}

func (s *searchEngines) has(engineType domain.SearchEngineType) bool {
	for _, engine := range *s.collectors {
		if engineType == engine.GetSearchEngineType() {
			return true
		}
	}
	return false
}
