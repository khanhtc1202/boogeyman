package data

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
