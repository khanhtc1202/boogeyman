package domain_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/internal/domain"
)

func TestQueryResultPool_FilterByEngineType(t *testing.T) {
	queryResultPool := fakePool()
	gQueryResult := queryResultPool.FilterByEngineType(domain.GOOGLE)
	if gQueryResult.Type() != domain.GOOGLE {
		t.Fatal("Fail test filter result from pool!")
	}
	if gQueryResult.GetQueryResults().Length() > 0 {
		t.Fatal("Fake data failed!")
	}
}

func TestQueryResultPool_GetSearchEngineList(t *testing.T) {
	queryResultPool := fakePool()
	engineList := queryResultPool.GetSearchEngineList()
	if engineList.Has(domain.GOOGLE) == false {
		t.Fatal("Fail test get search engine list! Missing type!")
	}
	if len(*engineList) != 3 {
		t.Fatal("Fail test get search engine list! Not enough type")
	}
}

func fakePool() *domain.QueryResultPool {
	pool := domain.EmptyQueryResultPool()
	pool.Add(domain.NewSearchEngine(domain.ASK, domain.EmptyQueryResult()))
	pool.Add(domain.NewSearchEngine(domain.BING, domain.EmptyQueryResult()))
	pool.Add(domain.NewSearchEngine(domain.GOOGLE, domain.EmptyQueryResult()))
	return pool
}
