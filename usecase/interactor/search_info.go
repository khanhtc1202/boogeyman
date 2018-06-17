package interactor

import (
	"github.com/khanhtc1202/boogeyman/config"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
	"github.com/pkg/errors"
)

type InfoSearch struct {
	ranker   *domain.Ranker
	poolRepo repository.QueryResultPool
}

func NewInfoSearch(
	poolRepo repository.QueryResultPool,
) *InfoSearch {
	return &InfoSearch{
		ranker:   domain.NewRanker(),
		poolRepo: poolRepo,
	}
}

func (i *InfoSearch) Search(
	queryString string,
	strategy domain.RankerStrategyType,
) (*domain.QueryResult, error) {
	resultPool, err := i.fetchData(domain.NewKeyword(queryString))
	if err != nil {
		return nil, errors.Wrap(err, "Error on fetch data from pool!\n")
	}
	switch strategy {
	case domain.TOP:
		return i.ranker.Top(resultPool)
	case domain.CROSS:
		return i.ranker.CrossMatch(resultPool)
	case domain.ALL:
		return i.ranker.All(resultPool,
			config.GetConfig().RankerConf.MaxReturnItems)
	default:
		return i.ranker.CrossMatch(resultPool)
	}
}

func (i *InfoSearch) fetchData(
	keyword *domain.Keyword,
) (*domain.QueryResultPool, error) {
	pool, err := i.poolRepo.FetchData(keyword)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching data from search engines!\n")
	}
	return pool, nil
}
