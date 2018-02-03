package interactor

import (
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
	"github.com/pkg/errors"
)

const MaxReturnItems = 5

type Ranker struct {
	materialPool repository.MaterialPool
}

func NewRanker(pool repository.MaterialPool) *Ranker {
	return &Ranker{
		materialPool: pool,
	}
}

func (r *Ranker) Top(searchEngines []search_engine.SearchEngineType) (*domain.ResultItems, error) {
	topResults := domain.EmptyResultItems()
	for _, searchEngine := range searchEngines {
		searchResult, err := r.materialPool.GetItemsBySearchEngine(searchEngine)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetch data from pool")
		}
		topResults.Add(searchResult.TopResult())
	}
	return topResults, nil
}

func (r *Ranker) CrossMatch(searchEngines []search_engine.SearchEngineType) (*domain.ResultItems, error) {
	topResults := domain.EmptyResultItems()
	// TODO implement return matched urls cross search engines
	return topResults, nil
}
