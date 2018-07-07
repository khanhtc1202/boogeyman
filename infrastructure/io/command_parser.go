package io

import (
	"flag"
	"github.com/khanhtc1202/boogeyman/infrastructure/io/data"
)

type CommandParse struct{}

func NewCommandParse() *CommandParse {
	return &CommandParse{}
}

func (c *CommandParse) ParseCommandParams() *data.CommandParams {
	var queryString string
	flag.StringVar(&queryString, "k", "boogeyman", "search (query) string")
	engine := flag.String("e", "all", "search engine(s): google | bing | ask | all")
	strategy := flag.String("s", "cross", "result show strategy: top | cross | all")

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show application version")
	flag.BoolVar(&showVersion, "version", false, "show application version")
	flag.Parse()

	return data.NewCommandParams(*engine, *strategy, queryString, showVersion)
}
