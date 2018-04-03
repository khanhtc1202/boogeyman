package main

import (
	"fmt"

	"strings"

	"github.com/khanhtc1202/boogeyman/adapter/controller"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/khanhtc1202/boogeyman/infrastructure/meta_info"
	spiderPool "github.com/khanhtc1202/boogeyman/infrastructure/service"
)

var (
	version   string
	revision  string
	buildDate string
	goVersion string
	mode      string
)

var metaInfo = meta_info.NewMetaInfo(
	version,
	revision,
	buildDate,
	goVersion,
	mode,
)

func main() {
	commandParse := controller.NewCommandParse()

	// check meta_info
	commandParse.ShowInfo(metaInfo)

	// parse command params
	cmdParams := commandParse.ParseCommandParams()

	materialPool := MaterialPoolFactory(cmdParams.Engine)
	materialPool.Fetch(domain.NewKeyword(cmdParams.QueryString))

	boogeyman := controller.NewBoogeyman(materialPool)

	results, err := boogeyman.ShowSearchResult(SetShowStrategy(cmdParams.Strategy), materialPool.GetSearchEngineList())
	if err != nil {
		fmt.Println("Error : ", err)
		panic("Error on show search results!")
	}
	for _, result := range *results {
		fmt.Println(result.Show())
	}
}

func MaterialPoolFactory(selectedEngine string) *repository.MaterialPool {
	collectors := service.EmptyCollectorList()
	switch strings.ToUpper(selectedEngine) {
	case search_engine.GOOGLE.String():
		collectors.Add(spiderPool.NewGoogleSpider())
		break
	case search_engine.BING.String():
		collectors.Add(spiderPool.NewBingSpider())
		break
	case search_engine.ASK.String():
		collectors.Add(spiderPool.NewAskSpider())
		break
	default:
		collectors.Add(spiderPool.NewAskSpider())
		collectors.Add(spiderPool.NewBingSpider())
		collectors.Add(spiderPool.NewGoogleSpider())
	}
	return repository.NewMaterialPool(*collectors)
}

func SetShowStrategy(selectedStrategy string) domain.StrategyType {
	switch strings.ToUpper(selectedStrategy) {
	case domain.TOP.String():
		return domain.TOP
	case domain.CROSS.String():
		return domain.CROSS
	default:
		return domain.ALL
	}
}
