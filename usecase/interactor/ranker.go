package interactor

import (
	"github.com/khanhtc1202/boogeyman/config"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
	"github.com/pkg/errors"
)

type Ranker struct {
	materialPool repository.MaterialPool
}

func NewRanker(pool repository.MaterialPool) *Ranker {
	return &Ranker{
		materialPool: pool,
	}
}

func (r *Ranker) Top(searchEngines *domain.SearchEngineList) (*domain.ResultItems, error) {
	topResults := domain.EmptyResultItems()
	for _, searchEngine := range *searchEngines {
		searchResult, err := r.materialPool.GetItemsFromSearchEngine(searchEngine)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetch data from pool")
		}
		topResults.Add(searchResult.TopResult())
	}
	return topResults, nil
}

func (r *Ranker) CrossMatch(searchEngines *domain.SearchEngineList) (*domain.ResultItems, error) {
	crossMatchedResults := domain.EmptyResultItems()
	for _, searchEngine := range *searchEngines {
		searchResult, err := r.materialPool.GetItemsFromSearchEngine(searchEngine)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetch data from pool")
		}
		crossMatchedResults.Concatenate(searchResult.GetResults())
	}
	return crossMatchedResults.DuplicateElements(), nil
}

func (r *Ranker) None(searchEngines *domain.SearchEngineList) (*domain.ResultItems, error) {
	allResults := domain.EmptyResultItems()
	for _, searchEngine := range *searchEngines {
		searchResult, err := r.materialPool.GetItemsFromSearchEngine(searchEngine)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetch data from pool")
		}
		allResults.Concatenate(searchResult.GetResults())
	}
	allResults.RemoveDuplicates()
	return allResults.Limit(config.GetConfig().RankerConf.MaxReturnItems), nil
}
