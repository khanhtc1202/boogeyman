package repository

import (
	"github.com/khanhtc1202/boogeyman/domain"
)

type MaterialPool interface {
	GetItemsFromSearchEngine(searchEngineType domain.SearchEngineType) (*domain.SearchEngine, error)
}
