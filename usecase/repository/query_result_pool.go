package repository

import (
	"github.com/khanhtc1202/boogeyman/domain"
)

type QueryResultPool interface {
	FetchData(keyword *domain.Keyword) (*domain.QueryResultPool, error)
}
