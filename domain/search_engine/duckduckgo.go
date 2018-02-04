package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type DuckDuckGo struct {
	Base
	keyword *domain.Keyword
	results *domain.ResultItems
}

func NewDuckDuckGo(keyword *domain.Keyword, results *domain.ResultItems) *DuckDuckGo {
	return &DuckDuckGo{
		keyword: keyword,
		results: results,
	}
}

func (d *DuckDuckGo) Type() SearchEngineType {
	return DUCKDUCKGO
}

func (d *DuckDuckGo) TopResult() *domain.ResultItem {
	return d.results.First()
}

func (d *DuckDuckGo) GetResults() *domain.ResultItems {
	return d.results
}
