package service

import (
	"github.com/khanhtc1202/boogeyman/domain"
)

type Collector interface {
	GetSearchEngineType() domain.SearchEngineType
	Query(keyword *domain.Keyword) (*domain.SearchEngine, error)
}

type CollectorList []Collector

func EmptyCollectorList() *CollectorList {
	return &CollectorList{}
}

func (c *CollectorList) Add(collector Collector) {
	*c = append(*c, collector)
}
