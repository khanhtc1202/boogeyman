package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type Bing struct {
	Base
	keyword string
	results *domain.ResultItems
}

func NewBing(keyword string, results *domain.ResultItems) *Bing {
	return &Bing{
		keyword: keyword,
		results: results,
	}
}

func (b *Bing) Type() SearchEngineType {
	return BING
}
