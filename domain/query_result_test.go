package domain_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/domain"
)

func TestResultItems_Add(t *testing.T) {
	resultItems := itemListFactory()

	item := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")
	resultItems.Add(item)

	if len(*resultItems) != 4 {
		t.Fatal("Fail test add item to list items!")
	}
}

func TestResultItems_RemoveDuplicates(t *testing.T) {
	resultItems := itemListFactory()

	resultItems.RemoveDuplicates()

	if len(*resultItems) != 2 {
		t.Fatal("Fail test remove duplicate item to list items!")
	}
}

func TestResultItems_DuplicateElements(t *testing.T) {
	resultItems := itemListFactory()

	if len(*resultItems.DuplicateElements()) != 1 {
		t.Fatal("Fail test remove duplicate item to list items!")
	}
}

func itemListFactory() *domain.QueryResult {
	resultItems := domain.EmptyQueryResult()

	item1 := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")
	item2 := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")
	item3 := domain.NewResultItem("dummy", "dummy", "dummy", "http://sample...")

	resultItems.Add(item1)
	resultItems.Add(item2)
	resultItems.Add(item3)
	return resultItems
}
