package interactor_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
)

type MaterialPoolMock struct {
	repository.MaterialPool
}

func (m *MaterialPoolMock) GetItemsBySearchEngine(searchEngineType search_engine.SearchEngineType) (search_engine.Base, error) {
	switch searchEngineType {
	case search_engine.GOOGLE:
		return search_engine.NewGoogle("key", fakeResultListSet1()), nil
	case search_engine.BING:
		return search_engine.NewBing("key", fakeResultListSet2()), nil
	case search_engine.DUCKDUCKGO:
		return search_engine.NewDuckDuckGo("key", fakeResultListSet3()), nil
	default:
		return nil, nil
	}
}

func TestRanker_CrossMatch(t *testing.T) {
	ranker := interactor.NewRanker(&MaterialPoolMock{})

	sEngineList := search_engine.EmptySearchEngineList()
	sEngineList.Add(search_engine.GOOGLE)
	sEngineList.Add(search_engine.BING)
	sEngineList.Add(search_engine.DUCKDUCKGO)

	results, err := ranker.CrossMatch(*sEngineList)
	if err != nil {
		t.Fatal("Fail running test get result urls by cross match")
	}
	if len(*results) != 1 {
		t.Fatal("Fail test logic get result urls by cross match")
	}
}

func TestRanker_Top(t *testing.T) {
	ranker := interactor.NewRanker(&MaterialPoolMock{})

	sEngineList := search_engine.EmptySearchEngineList()
	sEngineList.AddAll()

	results, err := ranker.Top(*sEngineList)
	if err != nil {
		t.Fatal("Fail running test get result urls by top ranking")
	}
	if len(*results) != len(*sEngineList) {
		t.Fatal("Fail test logic get result urls by top ranking")
	}
}

func TestRanker_None(t *testing.T) {
	ranker := interactor.NewRanker(&MaterialPoolMock{})

	sEngineList := search_engine.EmptySearchEngineList()
	sEngineList.AddAll()

	results, err := ranker.None(*sEngineList)
	if err != nil {
		t.Fatal("Fail running test show all result urls")
	}
	if len(*results) > interactor.MaxReturnItems {
		t.Fatal("Fail test logic show all result urls")
	}
}

func fakeResultListSet1() *domain.ResultItems {
	fakeResult1 := domain.NewResultItem("timestamp", "google 1", "my desc", "http://sample.com/acv")
	fakeResult2 := domain.NewResultItem("timestamp", "google 2", "my desc", "http://...")

	results := domain.EmptyResultItems()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}

func fakeResultListSet2() *domain.ResultItems {
	fakeResult1 := domain.NewResultItem("timestamp", "bing 1 ", "my desc", "http://...")
	fakeResult2 := domain.NewResultItem("timestamp", "bing 2", "my desc", "http://sample.com/123")

	results := domain.EmptyResultItems()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}

func fakeResultListSet3() *domain.ResultItems {
	fakeResult1 := domain.NewResultItem("timestamp", "duck 1 ", "my desc", "http://...")
	fakeResult2 := domain.NewResultItem("timestamp", "duck 2", "my desc", "http://...")

	results := domain.EmptyResultItems()
	results.Add(fakeResult1)
	results.Add(fakeResult2)
	return results
}
