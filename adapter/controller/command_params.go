package controller

import (
	"flag"
	"fmt"

	"os"

	"github.com/khanhtc1202/boogeyman/infrastructure/meta_info"
)

type CommandParams struct {
	Engine      string
	Strategy    string
	QueryString string
}

func NewCommandParams(
	engine string,
	strategy string,
	queryString string,
) *CommandParams {
	return &CommandParams{
		Engine:      engine,
		Strategy:    strategy,
		QueryString: queryString,
	}
}

type CommandParse struct{}

func NewCommandParse() *CommandParse {
	return &CommandParse{}
}

func (c *CommandParse) ParseCommandParams() *CommandParams {
	var queryString string
	flag.StringVar(&queryString, "k", "github.com/khanhtc1202/boogeyman", "search (query) string")
	engine := flag.String("e", "all", "search engine(s): google | bing | ask | all")
	strategy := flag.String("s", "all", "result show strategy: top | cross | all")

	flag.Parse()

	return NewCommandParams(*engine, *strategy, queryString)
}

func (c *CommandParse) ShowInfo(metaInfo *meta_info.MetaInfo) {
	var showVersion = false
	flag.BoolVar(&showVersion, "v", false, "show application version")
	flag.BoolVar(&showVersion, "version", false, "show application version")
	flag.Parse()

	if showVersion {
		fmt.Printf(metaInfo.GetMetaInfo())
		os.Exit(0)
	}
}
