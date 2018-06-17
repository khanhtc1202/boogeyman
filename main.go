package main

import (
	"fmt"

	"strings"

	"os"

	"github.com/khanhtc1202/boogeyman/adapter/controller"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/cross_cutting/io"
	"github.com/khanhtc1202/boogeyman/domain"
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

	// parse command params
	cmdParams := commandParse.ParseCommandParams()

	// check meta_info
	if cmdParams.ShowVersion {
		ShowMetaInfo(metaInfo)
	}

	resultPoolRepo := MaterialPoolFactory(cmdParams.Engine)
	boogeyman := controller.NewBoogeyman(resultPoolRepo)

	results, err := boogeyman.Search(cmdParams.QueryString, SetQueryStrategy(cmdParams.Strategy))
	if err != nil {
		io.Errorln(err)
		os.Exit(1)
	}
	for _, result := range *results {
		fmt.Println(result.Show())
	}
}

func ShowMetaInfo(metaInfo *meta_info.MetaInfo) {
	io.Infof(metaInfo.GetMetaInfo())
	os.Exit(0)
}

func MaterialPoolFactory(selectedEngine string) *repository.QueryResultPool {
	collectors := service.EmptyCollectorList()
	switch strings.ToUpper(selectedEngine) {
	case domain.GOOGLE.String():
		collectors.Add(spiderPool.NewGoogleSpider())
		break
	case domain.BING.String():
		collectors.Add(spiderPool.NewBingSpider())
		break
	case domain.ASK.String():
		collectors.Add(spiderPool.NewAskSpider())
		break
	default:
		collectors.Add(spiderPool.NewAskSpider())
		collectors.Add(spiderPool.NewBingSpider())
		collectors.Add(spiderPool.NewGoogleSpider())
	}
	return repository.NewResultPool(*collectors)
}

func SetQueryStrategy(selectedStrategy string) domain.RankerStrategyType {
	switch strings.ToUpper(selectedStrategy) {
	case domain.TOP.String():
		return domain.TOP
	case domain.CROSS.String():
		return domain.CROSS
	default:
		return domain.ALL
	}
}
