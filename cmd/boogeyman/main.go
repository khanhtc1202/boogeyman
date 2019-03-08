package main

import (
	"os"

	"strings"

	"github.com/khanhtc1202/boogeyman/internal/adapter/controller"
	"github.com/khanhtc1202/boogeyman/internal/adapter/presenter/console"
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/gateway/repository"
	"github.com/khanhtc1202/boogeyman/internal/gateway/service"
	"github.com/khanhtc1202/boogeyman/internal/infrastructure/cmd"
	"github.com/khanhtc1202/boogeyman/internal/infrastructure/meta_info"
	spiderPool "github.com/khanhtc1202/boogeyman/internal/infrastructure/service"
	"github.com/khanhtc1202/boogeyman/tools/io"
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
	commandParser := cmd.NewCommandParser()

	// parse command params
	cmdParams := commandParser.ParseParams()

	// check meta_info
	if cmdParams.ShowVersion {
		ShowMetaInfo(metaInfo)
	}

	resultPoolRepo := MaterialPoolFactory(cmdParams.Engine)
	textPresenter := console.NewColorfulTextPresenter()
	queryStrategy := SetQueryStrategy(cmdParams.Strategy)

	infoSearchCtl := controller.NewInfoSearch(textPresenter, resultPoolRepo)

	err := infoSearchCtl.Search(cmdParams.QueryString, queryStrategy)
	if err != nil {
		io.Errorln(err)
		os.Exit(1)
	}
}

func ShowMetaInfo(metaInfo *meta_info.MetaInfo) {
	io.Infof(metaInfo.GetMetaInfo())
	os.Exit(0)
}

func MaterialPoolFactory(selectedEngine string) *repository.SearchEnginePool {
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
	case domain.YAHOO.String():
		collectors.Add(spiderPool.NewYahooSpider())
		break
	default:
		collectors.Add(spiderPool.NewAskSpider())
		collectors.Add(spiderPool.NewBingSpider())
		collectors.Add(spiderPool.NewGoogleSpider())
		collectors.Add(spiderPool.NewYahooSpider())
	}
	return repository.NewSearchEnginePool(*collectors)
}

func SetQueryStrategy(selectedStrategy string) domain.FilterStrategyType {
	switch strings.ToUpper(selectedStrategy) {
	case domain.TOP.String():
		return domain.TOP
	case domain.CROSS.String():
		return domain.CROSS
	default:
		return domain.ALL
	}
}
