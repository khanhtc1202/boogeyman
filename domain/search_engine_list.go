package domain

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
	s.Add(ASK)
}

func (s *SearchEngineList) Has(searchEngineType SearchEngineType) bool {
	for _, engineType := range *s {
		if searchEngineType == engineType {
			return true
		}
	}
	return false
}
