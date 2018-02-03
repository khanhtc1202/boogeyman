package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type DuckDuckGo struct {
	Base
	keyword string
	results *domain.ResultItems
}

func NewDuckDuckGo(keyword string, results *domain.ResultItems) *DuckDuckGo {
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
