package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type SearchEngine interface {
	Type() SearchEngineType
	TopResult() *domain.ResultItem
	GetResults() *domain.ResultItems
}
