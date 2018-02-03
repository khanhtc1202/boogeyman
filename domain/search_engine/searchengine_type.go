package search_engine

type SearchEngineType int

const (
	GOOGLE SearchEngineType = iota
	BING
	DUCKDUCKGO
)

func (s SearchEngineType) String() string {
	switch s {
	case GOOGLE:
		return "GOOGLE"
	case BING:
		return "BING"
	case DUCKDUCKGO:
		return "DUCKDUCKGO"
	}
	return "Unknown"
}
