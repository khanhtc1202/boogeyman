package controller

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
)

type Boogeyman struct {
	interactor *interactor.Ranker
}

func NewBoogeyman(
	materialPool repository.MaterialPool,
) *Boogeyman {
	return &Boogeyman{
		interactor: interactor.NewRanker(materialPool),
	}
}

func (b *Boogeyman) ShowSearchResult(strategy domain.StrategyType, searchEngineList *search_engine.SearchEngineList) (*domain.ResultItems, error) {
	switch strategy {
	case domain.TOP:
		return b.interactor.Top(searchEngineList)
	case domain.CROSS:
		return b.interactor.CrossMatch(searchEngineList)
	case domain.ALL:
		return b.interactor.None(searchEngineList)
	default:
		return b.interactor.None(searchEngineList)
	}
}
