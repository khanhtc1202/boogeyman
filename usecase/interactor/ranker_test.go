package interactor_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/config"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
	"github.com/khanhtc1202/boogeyman/usecase/repository"
)

type MaterialPoolMock struct {
	repository.MaterialPool
}

func (m *MaterialPoolMock) GetItemsFromSearchEngine(searchEngineType domain.SearchEngineType) (*domain.SearchEngine, error) {
	switch searchEngineType {
	case domain.GOOGLE:
		return domain.NewSearchEngine(domain.GOOGLE, fakeResultListSet1()), nil
	case domain.BING:
		return domain.NewSearchEngine(domain.BING, fakeResultListSet2()), nil
	case domain.DUCKDUCKGO:
		return domain.NewSearchEngine(domain.DUCKDUCKGO, fakeResultListSet3()), nil
	default:
		return nil, nil
	}
}

func TestRanker_CrossMatch(t *testing.T) {
	ranker := interactor.NewRanker(&MaterialPoolMock{})
	sEngineList := fakeSearchEngineList()

	results, err := ranker.CrossMatch(sEngineList)
	if err != nil {
		t.Fatal("Fail running test get result urls by cross match")
	}
	if len(*results) != 1 {
		t.Fatal("Fail test logic get result urls by cross match")
	}
}

func TestRanker_Top(t *testing.T) {
	ranker := interactor.NewRanker(&MaterialPoolMock{})
	sEngineList := fakeSearchEngineList()

	results, err := ranker.Top(sEngineList)
	if err != nil {
		t.Fatal("Fail running test get result urls by top ranking")
	}
	if len(*results) != len(*sEngineList) {
		t.Fatal("Fail test logic get result urls by top ranking")
	}
}

func TestRanker_None(t *testing.T) {
	ranker := interactor.NewRanker(&MaterialPoolMock{})
	sEngineList := fakeSearchEngineList()

	results, err := ranker.None(sEngineList)
	if err != nil {
		t.Fatal("Fail running test show all result urls")
	}
	if len(*results) > config.GetConfig().RankerConf.MaxReturnItems {
		t.Fatal("Fail test logic show all result urls")
	}
}

func fakeSearchEngineList() *domain.SearchEngineList {
	sEngineList := domain.EmptySearchEngineList()
	sEngineList.Add(domain.GOOGLE)
	sEngineList.Add(domain.BING)
	sEngineList.Add(domain.DUCKDUCKGO)
	return sEngineList
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
