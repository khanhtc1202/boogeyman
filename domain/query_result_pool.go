package domain

type QueryResultPool []*SearchEngine

func EmptyResultPool() *QueryResultPool {
	return &QueryResultPool{}
}

func (r *QueryResultPool) Add(resultData *SearchEngine) {
	*r = append(*r, resultData)
}

func (r *QueryResultPool) FilterByEngineType(searchEngineType SearchEngineType) *SearchEngine {
	for _, resultData := range *r {
		if resultData.Type() == searchEngineType {
			return resultData
		}
	}
	return nil
}
