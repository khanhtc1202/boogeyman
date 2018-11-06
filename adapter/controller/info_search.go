package controller

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/usecase/presenter"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
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
	_, err := b.interactor.Search(queryString, strategy)
	if err != nil {
		return err
	}
	//b.interactor.PrintResults(queryResults)
	return nil
}
