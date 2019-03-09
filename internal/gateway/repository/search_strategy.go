package repository

import (
	"github.com/khanhtc1202/boogeyman/config"
	"github.com/khanhtc1202/boogeyman/internal/domain"
)

type searchStrategies struct {
}

func SearchStrategies() *searchStrategies {
	return &searchStrategies{}
}

func (s *searchStrategies) GetStrategyByType(
	fType domain.FilterStrategyType,
	engines *domain.SearchEnginePool,
) domain.FilterSearch {
	switch fType {
	case domain.MERGE:
		maxReturnItem := config.GetConfig().RankerConf.MaxReturnItems
		return domain.MergeResultsByEngines(engines, maxReturnItem)
	case domain.TOP:
		return domain.TopResultsByEngines(engines)
	case domain.CROSS:
		return domain.CrossMatchByEngines(engines)
	default:
		// return cross match strategy by default
		return domain.CrossMatchByEngines(engines)
	}
}
