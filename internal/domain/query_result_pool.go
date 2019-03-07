package domain

type QueryResultPool map[SearchEngineType]*SearchEngine

func EmptyQueryResultPool() *QueryResultPool {
	return &QueryResultPool{}
}

func (q *QueryResultPool) Add(queryResult *SearchEngine) {
	(*q)[queryResult.Type()] = queryResult
}

func (q *QueryResultPool) FilterByEngineType(
	searchEngineType SearchEngineType,
) *SearchEngine {
	return (*q)[searchEngineType]
}

func (q *QueryResultPool) GetSearchEngineList() *SearchEngineList {
	sEngineList := EmptySearchEngineList()
	for _, engine := range *q {
		sEngineList.Add(engine.Type())
	}
	return sEngineList
}
