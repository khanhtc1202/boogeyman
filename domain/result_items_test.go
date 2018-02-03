package domain_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/domain"
)

func TestResultItems_Add(t *testing.T) {
	resultItems := itemListFactory()

	item := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")
	resultItems.Add(item)

	if len(*resultItems) != 3 {
		t.Fatal("Fail test add item to list items!")
	}
}

func TestResultItems_RemoveDuplicates(t *testing.T) {
	resultItems := itemListFactory()

	resultItems.RemoveDuplicates()

	if len(*resultItems) != 1 {
		t.Fatal("Fail test remove duplicate item to list items!")
	}
}

func itemListFactory() *domain.ResultItems {
	resultItems := domain.EmptyResultItems()

	item1 := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")
	item2 := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")

	resultItems.Add(item1)
	resultItems.Add(item2)
	return resultItems
}
