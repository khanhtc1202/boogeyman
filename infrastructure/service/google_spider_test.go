package service_test

import (
	"testing"

	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/khanhtc1202/boogeyman/infrastructure/service"
)

func TestGoogleSpider_Query(t *testing.T) {
	keyword := domain.NewKeyword("sample")
	bingSpider := service.NewGoogleSpider()

	result, err := bingSpider.Query(keyword)
	if err != nil {
		t.Fatal("Fail test query data from search engine")
	}
	if len(*result.GetResults()) < 1 {
		t.Fatal("Fail test query data from se, maybe error on internet connection")
	}
	if result.Type() != search_engine.GOOGLE {
		t.Fatal("Fail test query data from se, error search engine type")
	}
}
