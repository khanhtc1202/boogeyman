package search_engine

type ResultDataPool []Base

func EmptyResultDataPool() *ResultDataPool {
	return &ResultDataPool{}
}

func (r *ResultDataPool) Add(resultData Base) {
	*r = append(*r, resultData)
}

func (r *ResultDataPool) FilterByEngineType(searchEngineType SearchEngineType) Base {
	for _, resultData := range *r {
		if resultData.Type() == searchEngineType {
			return resultData
		}
	}
	return nil
}
