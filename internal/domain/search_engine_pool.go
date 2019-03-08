package domain

type SearchEnginePool map[SearchEngineType]*SearchEngine

func EmptySearchEnginePool() *SearchEnginePool {
	return &SearchEnginePool{}
}

func (q *SearchEnginePool) Add(queryResult *SearchEngine) {
	(*q)[queryResult.Type()] = queryResult
}

func (q *SearchEnginePool) FilterByEngineType(
	searchEngineType SearchEngineType,
) *SearchEngine {
	return (*q)[searchEngineType]
}

func (q *SearchEnginePool) GetSearchEngineList() *SearchEngineTypeList {
	sEngineList := EmptySearchEngineTypeList()
	for _, engine := range *q {
		sEngineList.Add(engine.Type())
	}
	return sEngineList
}
