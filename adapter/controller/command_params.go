package controller

import "flag"

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
	flag.StringVar(&queryString, "keyword", "bar", "search (query) string")
	engine := flag.String("engine", "all", "search engine: google | bing | duckduckgo | all")
	strategy := flag.String("strategy", "all", "result show strategy: top | cross | all")

	flag.Parse()

	return NewCommandParams(*engine, *strategy, queryString)
}
