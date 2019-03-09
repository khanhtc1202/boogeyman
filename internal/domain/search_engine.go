package domain

import "strings"

/*
Search Engine Entity
*/
type SearchEngine struct {
	eType        SearchEngineType
	queryResults *QueryResults
}

func NewSearchEngine(
	eType SearchEngineType,
	results *QueryResults,
) *SearchEngine {
	return &SearchEngine{
		eType:        eType,
		queryResults: results,
	}
}

func (s *SearchEngine) Type() SearchEngineType {
	return s.eType
}

func (s *SearchEngine) TopResult() ComparableResultItem {
	return s.queryResults.First()
}

func (s *SearchEngine) GetQueryResults() *QueryResults {
	return s.queryResults
}

/*
Search Engine Type
*/
type SearchEngineType int

const (
	ALL SearchEngineType = iota
	GOOGLE
	BING
	DUCKDUCKGO
	ASK
	YANDEX
	YAHOO
)

func FactorySearchEngineType(sType string) SearchEngineType {
	switch strings.ToUpper(sType) {
	case GOOGLE.String():
		return GOOGLE
	case BING.String():
		return BING
	case ASK.String():
		return ASK
	case YAHOO.String():
		return YAHOO
	default:
		return ALL
	}
}

func (s SearchEngineType) String() string {
	switch s {
	case ALL:
		return "ALL"
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
	case YAHOO:
		return "YAHOO"
	}
	return "ALL"
}

/*
Search Engine Type Collection
*/
type SearchEngineTypeList []SearchEngineType

func EmptySearchEngineTypeList() *SearchEngineTypeList {
	return &SearchEngineTypeList{}
}

func (s *SearchEngineTypeList) Add(searchEngineType SearchEngineType) {
	*s = append(*s, searchEngineType)
}

func (s *SearchEngineTypeList) AddAll() {
	s.Add(GOOGLE)
	s.Add(BING)
	s.Add(ASK)
}

func (s *SearchEngineTypeList) Has(searchEngineType SearchEngineType) bool {
	for _, engineType := range *s {
		if searchEngineType == engineType {
			return true
		}
	}
	return false
}
