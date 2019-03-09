package domain_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/internal/domain"
)

func TestSearchStrategy_CrossMatch(t *testing.T) {
	crossFilter := domain.CrossMatchByEngines(fakeSearchEnginePool())

	results, err := crossFilter.Filter()
	if err != nil {
		t.Fatal("Fail running test get result urls by cross match")
	}
	if len(*results) != 1 {
		t.Fatal("Fail test logic get result urls by cross match")
	}
}

func TestSearchStrategy_TopResults(t *testing.T) {
	topFilter := domain.TopResultsByEngines(fakeSearchEnginePool())
	sEngineList := fakeSearchEngineTypeList()

	results, err := topFilter.Filter()
	if err != nil {
		t.Fatal("Fail running test get result urls by top ranking")
	}
	if len(*results) != len(*sEngineList) {
		t.Fatal("Fail test logic get result urls by top ranking")
	}
}

func TestSearchStrategy_TopOfEmptyQueryResults(t *testing.T) {
	topFilter := domain.TopResultsByEngines(fakeSearchEnginePoolEmptyCase())
	sEngineList := fakeSearchEngineTypeList()

	results, err := topFilter.Filter()
	if err != nil {
		t.Fatal("Fail running test get result urls by top ranking")
	}
	if len(*results) == len(*sEngineList) {
		t.Fatal("Fail test logic get result urls by top ranking")
	}
}

func TestSearchStrategy_MergeAllQueryResults(t *testing.T) {
	maxReturnItem := 20
	mergeFilter := domain.MergeResultsByEngines(fakeSearchEnginePool(), maxReturnItem)

	results, err := mergeFilter.Filter()
	if err != nil {
		t.Fatal("Fail running test show all result urls")
	}
	if len(*results) > maxReturnItem {
		t.Fatal("Fail test logic show all result urls")
	}
}

func fakeSearchEnginePool() *domain.SearchEnginePool {
	pool := domain.EmptySearchEnginePool()
	pool.Add(domain.NewSearchEngine(domain.GOOGLE, fakeResultListSet1()))
	pool.Add(domain.NewSearchEngine(domain.BING, fakeResultListSet2()))
	pool.Add(domain.NewSearchEngine(domain.DUCKDUCKGO, fakeResultListSet3()))
	return pool
}

func fakeSearchEnginePoolEmptyCase() *domain.SearchEnginePool {
	pool := domain.EmptySearchEnginePool()
	pool.Add(domain.NewSearchEngine(domain.GOOGLE, domain.EmptyQueryResult()))
	pool.Add(domain.NewSearchEngine(domain.BING, fakeResultListSet2()))
	pool.Add(domain.NewSearchEngine(domain.DUCKDUCKGO, fakeResultListSet3()))
	return pool
}

func fakeSearchEngineTypeList() *domain.SearchEngineTypeList {
	sEngineList := domain.EmptySearchEngineTypeList()
	sEngineList.Add(domain.GOOGLE)
	sEngineList.Add(domain.BING)
	sEngineList.Add(domain.DUCKDUCKGO)
	return sEngineList
}

func fakeResultListSet1() *domain.QueryResults {
	fakeResult1 := domain.NewResultItem("timestamp", "google 1", "my desc", "http://sample.com/acv")
	fakeResult2 := domain.NewResultItem("timestamp", "google 2", "my desc", "http://...")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}

func fakeResultListSet2() *domain.QueryResults {
	fakeResult1 := domain.NewResultItem("timestamp", "bing 1 ", "my desc", "http://...")
	fakeResult2 := domain.NewResultItem("timestamp", "bing 2", "my desc", "http://sample.com/123")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}

func fakeResultListSet3() *domain.QueryResults {
	fakeResult1 := domain.NewResultItem("timestamp", "duck 1 ", "my desc", "http://...")
	fakeResult2 := domain.NewResultItem("timestamp", "duck 2", "my desc", "http://...")

	results := domain.EmptyQueryResult()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}
