package cmd

import (
	"flag"

	"github.com/khanhtc1202/boogeyman/infrastructure/cmd/data"
)

type Parser struct{}

func NewCommandParser() *Parser {
	return &Parser{}
}

func (c *Parser) ParseParams() *data.CommandParams {
	var queryString string
	flag.StringVar(&queryString, "k", "boogeyman", "search (query) string")
	engine := flag.String("e", "all", "search engine(s): google | bing | ask | yahoo | all")
	strategy := flag.String("s", "top", "result show strategy: top | cross | all")

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show application version")
	flag.BoolVar(&showVersion, "version", false, "show application version")
	flag.Parse()

	return data.NewCommandParams(*engine, *strategy, queryString, showVersion)
}
