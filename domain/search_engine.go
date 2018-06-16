package domain

type SearchEngine struct {
	EType   SearchEngineType
	results *ResultItems
}

func NewSearchEngine(
	eType SearchEngineType,
	results *ResultItems,
) *SearchEngine {
	return &SearchEngine{
		EType:   eType,
		results: results,
	}
}

func (s *SearchEngine) Type() SearchEngineType {
	return s.EType
}

func (s *SearchEngine) TopResult() *ResultItem {
	return s.results.First()
}

func (s *SearchEngine) GetResults() *ResultItems {
	return s.results
}
