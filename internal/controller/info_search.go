package controller

import (
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/internal/usecase/presenter"
	"github.com/khanhtc1202/boogeyman/internal/usecase/repository"
)

type InfoSearch struct {
	interactor *interactor.InfoSearch
}

func NewInfoSearch(
	searchStrategiesRepo repository.SearchStrategies,
	searchEnginesRepo repository.SearchEngines,
	presenter presenter.TextPresenter,
) *InfoSearch {
	return &InfoSearch{
		interactor: interactor.NewInfoSearch(
			searchStrategiesRepo,
			searchEnginesRepo,
			presenter),
	}
}

func (b *InfoSearch) Search(
	queryString string,
	engine string,
	strategy string,
) error {
	// adapt universe type (string) to internal type (domain type)
	keyword := domain.NewKeyword(queryString)

	searchEngineType := domain.FactorySearchEngineType(engine)

	strategyType := domain.FactoryFilterStrategyType(strategy)

	// call interactor
	return b.interactor.Search(keyword, searchEngineType, strategyType)
}
