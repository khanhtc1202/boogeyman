package main

import (
	"flag"
	"os"

	"github.com/khanhtc1202/boogeyman/internal/controller"
	"github.com/khanhtc1202/boogeyman/internal/gateway/repository"
	"github.com/khanhtc1202/boogeyman/internal/gateway/service"
	"github.com/khanhtc1202/boogeyman/internal/presenter/console"
	"github.com/khanhtc1202/boogeyman/pkg/io"
	"github.com/khanhtc1202/boogeyman/pkg/meta_info"
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

type commandParams struct {
	Engine      string
	Strategy    string
	QueryString string
	ShowVersion bool
}

func CommandParams(
	engine string,
	strategy string,
	queryString string,
	showVersion bool,
) *commandParams {
	return &commandParams{
		Engine:      engine,
		Strategy:    strategy,
		QueryString: queryString,
		ShowVersion: showVersion,
	}
}

func ParseParams() *commandParams {
	var queryString string
	flag.StringVar(&queryString, "k", "boogeyman", "search (query) string")
	engine := flag.String("e", "all", "search engine(s): google | bing | ask | yahoo | all")
	strategy := flag.String("s", "top", "result show strategy: top | cross | all")

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show application version")
	flag.BoolVar(&showVersion, "version", false, "show application version")
	flag.Parse()

	return CommandParams(*engine, *strategy, queryString, showVersion)
}

func main() {
	// parse command params
	cmdParams := ParseParams()

	// check meta_info
	if cmdParams.ShowVersion {
		ShowMetaInfo(metaInfo)
	}

	searchStrategiesRepo := repository.SearchStrategies()
	searchEnginesRepo := repository.SearchEngines(service.EmptyCollectorList())
	textPresenter := console.NewColorfulTextPresenter()

	infoSearchCtl := controller.NewInfoSearch(searchStrategiesRepo, searchEnginesRepo, textPresenter)

	err := infoSearchCtl.Search(cmdParams.QueryString, cmdParams.Engine, cmdParams.Strategy)
	if err != nil {
		io.Errorln(err)
		os.Exit(1)
	}
}

func ShowMetaInfo(metaInfo *meta_info.MetaInfo) {
	io.Infof(metaInfo.GetMetaInfo())
	os.Exit(0)
}
