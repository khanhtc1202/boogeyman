package search_engine

import "github.com/khanhtc1202/boogeyman/domain"

type Base interface {
	Type() SearchEngineType
	TopResult() *domain.ResultItem
}
