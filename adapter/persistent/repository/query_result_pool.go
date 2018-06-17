package repository

import (
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/pkg/errors"
)

type QueryResultPool struct {
	collectors []service.Collector
}

func NewResultPool(
	services []service.Collector,
) *QueryResultPool {
	return &QueryResultPool{
		collectors: services,
	}
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
