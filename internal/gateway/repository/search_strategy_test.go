package repository_test

import (
	"reflect"
	"testing"

	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/gateway/repository"
)

func TestSearchStrategies_GetStrategyByType_ReturnCrossMatchAsDefault(t *testing.T) {
	strategiesRepo := repository.SearchStrategies()

	strategy := strategiesRepo.GetStrategyByType(domain.UNKNOWN_STRATEGY, nil)
	if reflect.TypeOf(strategy) != reflect.TypeOf(domain.CrossMatchByEngines(nil)) {
		t.Fatal("Error not return cross match strategy by default")
	}
}
