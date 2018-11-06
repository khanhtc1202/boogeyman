package interactor

import (
	"github.com/khanhtc1202/boogeyman/config"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/presenter"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
	"github.com/pkg/errors"
)

type InfoSearch struct {
	ranker    *domain.Ranker
	poolRepo  repository.QueryResultPool
	presenter presenter.TextPresenter
}

func NewInfoSearch(
	presenter presenter.TextPresenter,
	poolRepo repository.QueryResultPool,
) *InfoSearch {
	return &InfoSearch{
		ranker:    domain.NewRanker(),
		poolRepo:  poolRepo,
		presenter: presenter,
	}
}

func (i *InfoSearch) Search(
	queryString string,
	strategy domain.RankerStrategyType,
) (*domain.QueryResult, error) {
	// fetch data from search engines
	resultPool, err := i.fetchData(domain.NewKeyword(queryString))
	if err != nil {
		return nil, errors.Wrap(err, "Error on fetch data from pool!\n")
	}

	// merge by strategy
	var queryResult *domain.QueryResult
	switch strategy {
	case domain.TOP:
		queryResult, err = i.ranker.Top(resultPool)
		break
	case domain.CROSS:
		queryResult, err = i.ranker.CrossMatch(resultPool)
		break
	case domain.ALL:
		queryResult, err = i.ranker.All(resultPool,
			config.GetConfig().RankerConf.MaxReturnItems)
		break
	default:
		queryResult, err = i.ranker.CrossMatch(resultPool)
		break
	}
	if err != nil {
		return nil, err
	}

	// printout
	if err = i.presenter.PrintList(queryResult); err != nil {
		return nil, err
	}

	return queryResult, nil
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
