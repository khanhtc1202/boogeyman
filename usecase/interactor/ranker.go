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

func (r *Ranker) Top(searchEngines *domain.SearchEngineList) (*domain.QueryResult, error) {
	topResults := domain.EmptyQueryResult()
	for _, searchEngine := range *searchEngines {
		searchResult, err := r.materialPool.GetItemsFromSearchEngine(searchEngine)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetch data from pool")
		}
		topResults.Add(searchResult.TopResult())
	}
	return topResults, nil
}

func (r *Ranker) CrossMatch(searchEngines *domain.SearchEngineList) (*domain.QueryResult, error) {
	crossMatchedResults := domain.EmptyQueryResult()
	for _, searchEngine := range *searchEngines {
		searchResult, err := r.materialPool.GetItemsFromSearchEngine(searchEngine)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetch data from pool")
		}
		crossMatchedResults.Concatenate(searchResult.GetQueryResults())
	}
	return crossMatchedResults.DuplicateElements(), nil
}

func (r *Ranker) None(searchEngines *domain.SearchEngineList) (*domain.QueryResult, error) {
	allResults := domain.EmptyQueryResult()
	for _, searchEngine := range *searchEngines {
		searchResult, err := r.materialPool.GetItemsFromSearchEngine(searchEngine)
		if err != nil {
			return nil, errors.Wrap(err, "Error on fetch data from pool")
		}
		allResults.Concatenate(searchResult.GetQueryResults())
	}
	allResults.RemoveDuplicates()
	return allResults.Limit(config.GetConfig().RankerConf.MaxReturnItems), nil
}
