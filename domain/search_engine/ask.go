package search_engine

import (
	"github.com/khanhtc1202/boogeyman/domain"
)

type Ask struct {
	Base
	keyword *domain.Keyword
	results *domain.ResultItems
}

func NewAsk(keyword *domain.Keyword, results *domain.ResultItems) *Ask {
	return &Ask{
		keyword: keyword,
		results: results,
	}
}

func (a *Ask) Type() SearchEngineType {
	return ASK
}

func (a *Ask) TopResult() *domain.ResultItem {
	return a.results.First()
}

func (a *Ask) GetResults() *domain.ResultItems {
	return a.results
}
