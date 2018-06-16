package domain

type QueryResult []*ResultItem

func EmptyQueryResult() *QueryResult {
	return &QueryResult{}
}

func (r *QueryResult) Add(resultItem *ResultItem) {
	if resultItem == nil {
		return
	}
	*r = append(*r, resultItem)
}

func (r *QueryResult) Concatenate(itemList *QueryResult) {
	for _, item := range *itemList {
		r.Add(item)
	}
}

func (r *QueryResult) First() *ResultItem {
	return (*r)[0]
}

func (r *QueryResult) RemoveDuplicates() {
	keys := make(map[string]bool)
	list := EmptyQueryResult()
	for _, entry := range *r {
		if _, value := keys[entry.url]; !value {
			keys[entry.url] = true
			list.Add(entry)
		}
	}
	*r = *list
}

func (r *QueryResult) DuplicateElements() *QueryResult {
	duplicateElements := EmptyQueryResult()

	keys := make(map[string]bool)
	for _, entry := range *r {
		if _, value := keys[entry.url]; !value {
			keys[entry.url] = true
		} else {
			duplicateElements.Add(entry)
		}
	}
	duplicateElements.RemoveDuplicates()
	return duplicateElements
}

func (r *QueryResult) Limit(limitSize int) *QueryResult {
	splitSlide := EmptyQueryResult()

	if limitSize > len(*r) {
		return r
	} else {
		for _, item := range (*r)[0:limitSize] {
			splitSlide.Add(item)
		}
		return splitSlide
	}
}
