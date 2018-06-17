package domain

type QueryResultPool []*SearchEngine

func EmptyQueryResultPool() *QueryResultPool {
	return &QueryResultPool{}
}

func (q *QueryResultPool) Add(queryResult *SearchEngine) {
	*q = append(*q, queryResult)
}

func (q *QueryResultPool) FilterByEngineType(searchEngineType SearchEngineType) *SearchEngine {
	for _, resultData := range *q {
		if resultData.Type() == searchEngineType {
			return resultData
		}
	}
	return nil
}

func (q *QueryResultPool) GetSearchEngineList() *SearchEngineList {
	// TODO change pool from array to map ???
	sEngineList := EmptySearchEngineList()
	for _, engine := range *q {
		sEngineList.Add(engine.EType)
	}
	return sEngineList
}
