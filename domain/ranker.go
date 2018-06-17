package domain

type Ranker struct{}

func NewRanker() *Ranker {
	return &Ranker{}
}

func (r *Ranker) Top(pool *QueryResultPool) (*QueryResult, error) {
	topResults := EmptyQueryResult()
	for _, searchEngine := range *pool.GetSearchEngineList() {
		searchEngine := pool.FilterByEngineType(searchEngine)
		topResults.Add(searchEngine.TopResult())
	}
	return topResults, nil
}

func (r *Ranker) CrossMatch(pool *QueryResultPool) (*QueryResult, error) {
	crossMatchedResults := EmptyQueryResult()
	for _, searchEngine := range *pool.GetSearchEngineList() {
		searchResult := pool.FilterByEngineType(searchEngine)
		crossMatchedResults.Concatenate(searchResult.GetQueryResults())
	}
	return crossMatchedResults.DuplicateElements(), nil
}

func (r *Ranker) All(pool *QueryResultPool, maxReturnItems int) (*QueryResult, error) {
	allResults := EmptyQueryResult()
	for _, searchEngine := range *pool.GetSearchEngineList() {
		searchResult := pool.FilterByEngineType(searchEngine)
		allResults.Concatenate(searchResult.GetQueryResults())
	}
	allResults.RemoveDuplicates()
	return allResults.Limit(maxReturnItems), nil
}
