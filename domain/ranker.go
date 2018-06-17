package domain

import (
	"github.com/khanhtc1202/boogeyman/config"
)

type Ranker struct {
	resultPool *QueryResultPool
}

func NewRanker(pool *QueryResultPool) *Ranker {
	return &Ranker{
		resultPool: pool,
	}
}

func (r *Ranker) Top() (*QueryResult, error) {
	topResults := EmptyQueryResult()
	for _, searchEngine := range *r.resultPool.GetSearchEngineList() {
		searchEngine := r.resultPool.FilterByEngineType(searchEngine)
		topResults.Add(searchEngine.TopResult())
	}
	return topResults, nil
}

func (r *Ranker) CrossMatch() (*QueryResult, error) {
	crossMatchedResults := EmptyQueryResult()
	for _, searchEngine := range *r.resultPool.GetSearchEngineList() {
		searchResult := r.resultPool.FilterByEngineType(searchEngine)
		crossMatchedResults.Concatenate(searchResult.GetQueryResults())
	}
	return crossMatchedResults.DuplicateElements(), nil
}

func (r *Ranker) None() (*QueryResult, error) {
	allResults := EmptyQueryResult()
	for _, searchEngine := range *r.resultPool.GetSearchEngineList() {
		searchResult := r.resultPool.FilterByEngineType(searchEngine)
		allResults.Concatenate(searchResult.GetQueryResults())
	}
	allResults.RemoveDuplicates()
	return allResults.Limit(config.GetConfig().RankerConf.MaxReturnItems), nil
}
