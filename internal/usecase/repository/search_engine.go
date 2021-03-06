package repository

import (
	"github.com/khanhtc1202/boogeyman/internal/domain"
)

type SearchEngines interface {
	FetchData(keyword domain.Keyword) (*domain.SearchEnginePool, error)
	AddEnginesByType(engineType domain.SearchEngineType) error
}
