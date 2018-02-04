package controller

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
)

type Boogeyman struct {
	searchEngineList *search_engine.SearchEngineList
	interactor       *interactor.Ranker
}

func NewBoogeyman(
	materialPool repository.MaterialPool,
	searchEngineList *search_engine.SearchEngineList,
) *Boogeyman {
	return &Boogeyman{
		interactor:       interactor.NewRanker(materialPool),
		searchEngineList: searchEngineList,
	}
}

func (b *Boogeyman) ShowSearchResult(strategy domain.StrategyType) (*domain.ResultItems, error) {
	switch strategy {
	case domain.TOP:
		return b.interactor.Top(b.searchEngineList)
	case domain.CROSS:
		return b.interactor.CrossMatch(b.searchEngineList)
	case domain.ALL:
		return b.interactor.None(b.searchEngineList)
	default:
		return b.interactor.None(b.searchEngineList)
	}
}
