package repository

import "github.com/khanhtc1202/boogeyman/internal/domain"

type SearchStrategies interface {
	GetStrategyByType(fType domain.FilterStrategyType, engines *domain.SearchEnginePool) domain.FilterSearch
}
