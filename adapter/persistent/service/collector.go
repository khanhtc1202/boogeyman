package service

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
)

type Collector interface {
	GetSearchEngineType() search_engine.SearchEngineType
	Query(keyword *domain.Keyword) (search_engine.Base, error)
}
