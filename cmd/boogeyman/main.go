package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"

	"github.com/khanhtc1202/boogeyman/internal/controller"
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/gateway/repository"
	"github.com/khanhtc1202/boogeyman/internal/gateway/service"
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
	textPresenter := NewColorfulTextPresenter()

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

type TextPresenter struct {
	writer io.UI
}

func NewColorfulTextPresenter() *TextPresenter {
	return &TextPresenter{
		writer: io.ColorfulConsole(),
	}
}

func (t *TextPresenter) PrintList(results *domain.QueryResults) error {
	for _, result := range *results {
		switch result.(type) {
		case *domain.UrlBaseResultItem:
			t.presentUrlBaseItem(result.(*domain.UrlBaseResultItem))
			continue
		default:
			return errors.New("Error not found presenter for this type of ResultItem")
		}
	}

	t.writer.Printf(color.HiCyanString("\nTotal %v result(s) founded!\n", len(*results)))
	return nil
}

func (t *TextPresenter) presentUrlBaseItem(result *domain.UrlBaseResultItem) {
	t.writer.Printf(color.HiGreenString("Title: %v \n", result.GetTitleString()))
	t.writer.Printf(color.YellowString("URL: %v \n", result.GetUrl()))
	t.writer.Printf(color.RedString("Description: ") + result.GetDescription() + "\n")
	t.writer.Printf(color.BlueString("Create At: %v \n", result.Time()))
	t.writer.Println("---------------------")
}
