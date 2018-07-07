package controller

import (
	"github.com/khanhtc1202/boogeyman/cross_cutting/common"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
)

type Boogeyman struct {
	interactor *interactor.InfoSearch
}

func NewBoogeyman(
	container common.IDIContainer,
) *Boogeyman {
	return &Boogeyman{
		interactor: container.SearchInfo(),
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
	b.interactor.PrintResults(queryResults)
	return nil
}
