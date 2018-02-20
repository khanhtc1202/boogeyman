package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type Yandex struct {
	Base
	keyword *domain.Keyword
	results *domain.ResultItems
}

func NewYandex(keyword *domain.Keyword, results *domain.ResultItems) *Yandex {
	return &Yandex{
		keyword: keyword,
		results: results,
	}
}

func (y *Yandex) Type() SearchEngineType {
	return YANDEX
}

func (y *Yandex) TopResult() *domain.ResultItem {
	return y.results.First()
}

func (y *Yandex) GetResults() *domain.ResultItems {
	return y.results
}
