package interactor

import (
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/usecase/presenter"
	"github.com/khanhtc1202/boogeyman/internal/usecase/repository"
	"github.com/pkg/errors"
)

type InfoSearch struct {
	searchStrategies repository.SearchStrategies
	searchEngines    repository.SearchEngines
	presenter        presenter.TextPresenter
}

func NewInfoSearch(
	strategiesRepo repository.SearchStrategies,
	searchEngines repository.SearchEngines,
	presenter presenter.TextPresenter,
) *InfoSearch {
	return &InfoSearch{
		searchStrategies: strategiesRepo,
		searchEngines:    searchEngines,
		presenter:        presenter,
	}
}

func (i *InfoSearch) Search(
	query domain.Keyword,
	strategy domain.FilterStrategyType,
) error {
	// fetch data from search engines
	engines, err := i.searchEngines.FetchData(query)
	if err != nil {
		return errors.Wrap(err, "Error on fetch data from pool!\n")
	}

	// filter result
	searchStrategy := i.searchStrategies.GetStrategyByType(strategy, engines)
	queryResult, err := searchStrategy.Filter()
	if err != nil {
		return errors.Wrap(err, "Error on filter results!\n")
	}

	// out
	if err = i.presenter.PrintList(queryResult); err != nil {
		return errors.Wrap(err, "Error on push results!\n")
	}

	return nil
}
