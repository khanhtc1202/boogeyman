package search_engine

type ResultDataPool []SearchEngine

func EmptyResultDataPool() *ResultDataPool {
	return &ResultDataPool{}
}

func (r *ResultDataPool) Add(resultData SearchEngine) {
	*r = append(*r, resultData)
}

func (r *ResultDataPool) FilterByEngineType(searchEngineType SearchEngineType) SearchEngine {
	for _, resultData := range *r {
		if resultData.Type() == searchEngineType {
			return resultData
		}
	}
	return nil
}
