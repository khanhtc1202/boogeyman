package domain

type SearchEngine struct {
	eType        SearchEngineType
	queryResults *QueryResult
}

func NewSearchEngine(
	eType SearchEngineType,
	results *QueryResult,
) *SearchEngine {
	return &SearchEngine{
		eType:        eType,
		queryResults: results,
	}
}

func (s *SearchEngine) Type() SearchEngineType {
	return s.eType
}

func (s *SearchEngine) TopResult() *ResultItem {
	return s.queryResults.First()
}

func (s *SearchEngine) GetQueryResults() *QueryResult {
	return s.queryResults
}
