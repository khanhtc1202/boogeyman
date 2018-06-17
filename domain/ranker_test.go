package domain_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/domain"
)

func TestRanker_CrossMatch(t *testing.T) {
	ranker := domain.NewRanker()

	results, err := ranker.CrossMatch(fakeQueryResultPoll())
	if err != nil {
		t.Fatal("Fail running test get result urls by cross match")
	}
	if len(*results) != 1 {
		t.Fatal("Fail test logic get result urls by cross match")
	}
}

func TestRanker_Top(t *testing.T) {
	ranker := domain.NewRanker()
	sEngineList := fakeSearchEngineList()

	results, err := ranker.Top(fakeQueryResultPoll())
	if err != nil {
		t.Fatal("Fail running test get result urls by top ranking")
	}
	if len(*results) != len(*sEngineList) {
		t.Fatal("Fail test logic get result urls by top ranking")
	}
}

func TestRanker_None(t *testing.T) {
	ranker := domain.NewRanker()
	maxReturnItem := 20

	results, err := ranker.All(fakeQueryResultPoll(), maxReturnItem)
	if err != nil {
		t.Fatal("Fail running test show all result urls")
	}
	if len(*results) > maxReturnItem {
		t.Fatal("Fail test logic show all result urls")
	}
}

func fakeQueryResultPoll() *domain.QueryResultPool {
	pool := domain.EmptyQueryResultPool()
	pool.Add(domain.NewSearchEngine(domain.GOOGLE, fakeResultListSet1()))
	pool.Add(domain.NewSearchEngine(domain.BING, fakeResultListSet2()))
	pool.Add(domain.NewSearchEngine(domain.DUCKDUCKGO, fakeResultListSet3()))
	return pool
}

func fakeSearchEngineList() *domain.SearchEngineList {
	sEngineList := domain.EmptySearchEngineList()
	sEngineList.Add(domain.GOOGLE)
	sEngineList.Add(domain.BING)
	sEngineList.Add(domain.DUCKDUCKGO)
	return sEngineList
}

func fakeResultListSet1() *domain.QueryResult {
	fakeResult1 := domain.NewResultItem("timestamp", "google 1", "my desc", "http://sample.com/acv")
	fakeResult2 := domain.NewResultItem("timestamp", "google 2", "my desc", "http://...")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}

func fakeResultListSet2() *domain.QueryResult {
	fakeResult1 := domain.NewResultItem("timestamp", "bing 1 ", "my desc", "http://...")
	fakeResult2 := domain.NewResultItem("timestamp", "bing 2", "my desc", "http://sample.com/123")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}

func fakeResultListSet3() *domain.QueryResult {
	fakeResult1 := domain.NewResultItem("timestamp", "duck 1 ", "my desc", "http://...")
	fakeResult2 := domain.NewResultItem("timestamp", "duck 2", "my desc", "http://...")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}
