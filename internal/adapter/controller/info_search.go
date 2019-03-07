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
	presenter presenter.TextPresenter,
	resultPoolRepo repository.QueryResultPool,
) *InfoSearch {
	return &InfoSearch{
		interactor: interactor.NewInfoSearch(presenter, resultPoolRepo),
	}
}

func (b *InfoSearch) Search(
	queryString string,
	strategy domain.RankerStrategyType,
) error {
	return b.interactor.Search(queryString, strategy)
}
