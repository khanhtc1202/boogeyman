package domain_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/domain"
)

func TestResultItems_Add(t *testing.T) {
	resultItems := domain.EmptyResultItems()

	item1 := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")
	item2 := domain.NewResultItem("dummy", "dummy", "dummy", "http://...")

	resultItems.Add(item1)
	resultItems.Add(item2)

	if len(*resultItems) != 2 {
		t.Fatal("Fail test add item to list items!")
	}
}
