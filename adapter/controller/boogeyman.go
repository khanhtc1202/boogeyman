package controller

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
	"github.com/pkg/errors"
)

type Boogeyman struct {
	interactor *interactor.InfoSearch
}

func NewBoogeyman(
	resultPoolRepo repository.QueryResultPool,
) *Boogeyman {
	return &Boogeyman{
		interactor: interactor.NewInfoSearch(resultPoolRepo),
	}
}

func (b *Boogeyman) Search(
	queryString string,
	strategy domain.RankerStrategyType,
) (*domain.QueryResult, error) {
	queryResults, err := b.interactor.Search(queryString, strategy)
	if err != nil {
		return nil, errors.Wrap(err, "Error on search keyword!\n")
	}
	return queryResults, nil
}
