package controller

import (
	"flag"
)

type CommandParams struct {
	Engine      string
	Strategy    string
	QueryString string
	ShowVersion bool
}

func NewCommandParams(
	engine string,
	strategy string,
	queryString string,
	showVersion bool,
) *CommandParams {
	return &CommandParams{
		Engine:      engine,
		Strategy:    strategy,
		QueryString: queryString,
		ShowVersion: showVersion,
	}
}

type CommandParse struct{}

func NewCommandParse() *CommandParse {
	return &CommandParse{}
}

func (c *CommandParse) ParseCommandParams() *CommandParams {
	var queryString string
	flag.StringVar(&queryString, "k", "boogeyman", "search (query) string")
	engine := flag.String("e", "all", "search engine(s): google | bing | ask | all")
	strategy := flag.String("s", "cross", "result show strategy: top | cross | all")

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show application version")
	flag.BoolVar(&showVersion, "version", false, "show application version")
	flag.Parse()

	return NewCommandParams(*engine, *strategy, queryString, showVersion)
}
