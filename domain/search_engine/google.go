package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type Google struct {
	SearchEngine
	keyword *domain.Keyword
	results *domain.ResultItems
}

func NewGoogle(keyword *domain.Keyword, results *domain.ResultItems) *Google {
	return &Google{
		keyword: keyword,
		results: results,
	}
}

func (g *Google) Type() SearchEngineType {
	return GOOGLE
}

func (g *Google) TopResult() *domain.ResultItem {
	return g.results.First()
}

func (g *Google) GetResults() *domain.ResultItems {
	return g.results
}
