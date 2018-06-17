package domain

type SearchEngine struct {
	EType        SearchEngineType
	queryResults *QueryResult
}

func NewSearchEngine(
	eType SearchEngineType,
	results *QueryResult,
) *SearchEngine {
	return &SearchEngine{
		EType:        eType,
		queryResults: results,
	}
}

func (s *SearchEngine) Type() SearchEngineType {
	return s.EType
}

func (s *SearchEngine) TopResult() *ResultItem {
	return s.queryResults.First()
}

func (s *SearchEngine) GetQueryResults() *QueryResult {
	return s.queryResults
}
