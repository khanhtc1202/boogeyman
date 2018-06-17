package domain

type SearchEngineType int

const (
	GOOGLE SearchEngineType = iota
	BING
	DUCKDUCKGO
	ASK
	YANDEX
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
	case YANDEX:
		return "YANDEX"
	}
	return "Unknown"
}
