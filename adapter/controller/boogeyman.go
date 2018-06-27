package controller

import (
	"github.com/khanhtc1202/boogeyman/adapter/controller/presenter"
	"github.com/khanhtc1202/boogeyman/adapter/presenter/console"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
)

type Boogeyman struct {
	interactor *interactor.InfoSearch
	presenter  presenter.TextPresenter
}

func NewBoogeyman(
	resultPoolRepo repository.QueryResultPool,
) *Boogeyman {
	return &Boogeyman{
		interactor: interactor.NewInfoSearch(resultPoolRepo),
		presenter:  console.NewColorfulTextPresenter(),
	}
}

func (b *Boogeyman) Search(
	queryString string,
	strategy domain.RankerStrategyType,
) error {
	queryResults, err := b.interactor.Search(queryString, strategy)
	if err != nil {
		return err
	}
	b.presenter.PrintList(queryResults)
	return nil
}
