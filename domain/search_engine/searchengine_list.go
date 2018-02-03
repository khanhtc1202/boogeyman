package search_engine

type SearchEngineList []SearchEngineType

func EmptySearchEngineList() *SearchEngineList {
	return &SearchEngineList{}
}

func (s *SearchEngineList) Add(searchEngineType SearchEngineType) {
	*s = append(*s, searchEngineType)
}

func (s *SearchEngineList) AddAll() {
	s.Add(GOOGLE)
	s.Add(BING)
	s.Add(DUCKDUCKGO)
}
