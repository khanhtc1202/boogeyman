package domain

/*
Ranker Strategy Type
*/
type RankerStrategyType int

const (
	ALL RankerStrategyType = iota
	TOP
	CROSS
)

func (s RankerStrategyType) String() string {
	switch s {
	case ALL:
		return "ALL"
	case CROSS:
		return "CROSS"
	case TOP:
		return "TOP"
	}
	return "ALL"
}

// TODO not to Universe
type Ranker struct{}

func NewRanker() *Ranker {
	return &Ranker{}
}

func (r *Ranker) Top(pool *SearchEnginePool) (*QueryResults, error) {
	topResults := EmptyQueryResult()
	for _, searchEngine := range *pool {
		if searchEngine.GetQueryResults().Length() > 0 {
			topResults.Add(searchEngine.TopResult())
		}
	}
	return topResults, nil
}

func (r *Ranker) CrossMatch(pool *SearchEnginePool) (*QueryResults, error) {
	crossMatchedResults := EmptyQueryResult()
	for _, searchEngine := range *pool {
		crossMatchedResults.Concatenate(searchEngine.GetQueryResults())
	}
	return crossMatchedResults.DuplicateElements(), nil
}

func (r *Ranker) All(pool *SearchEnginePool, maxReturnItems int) (*QueryResults, error) {
	allResults := EmptyQueryResult()
	for _, searchEngine := range *pool {
		allResults.Concatenate(searchEngine.GetQueryResults())
	}
	allResults.RemoveDuplicates()
	return allResults.Limit(maxReturnItems), nil
}
