package search_engine

type SearchEngineType int

const (
	GOOGLE SearchEngineType = iota
	BING
	DUCKDUCKGO
	ASK
)

func (s SearchEngineType) String() string {
	switch s {
	case GOOGLE:
		return "GOOGLE"
	case BING:
		return "BING"
	case DUCKDUCKGO:
		return "DUCKDUCKGO"
	case ASK:
		return "ASK"
	}
	return "Unknown"
}
