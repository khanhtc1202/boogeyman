package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type Bing struct {
	SearchEngine
	keyword *domain.Keyword
	results *domain.ResultItems
}

func NewBing(keyword *domain.Keyword, results *domain.ResultItems) *Bing {
	return &Bing{
		keyword: keyword,
		results: results,
	}
}

func (b *Bing) Type() SearchEngineType {
	return BING
}

func (b *Bing) TopResult() *domain.ResultItem {
	return b.results.First()
}

func (b *Bing) GetResults() *domain.ResultItems {
	return b.results
}
