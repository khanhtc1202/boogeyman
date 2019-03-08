package domain

type QueryResults []ComparableResultItem

func EmptyQueryResult() *QueryResults {
	return &QueryResults{}
}

func (r *QueryResults) Add(resultItem ComparableResultItem) {
	if resultItem == nil {
		return
	}
	*r = append(*r, resultItem)
}

func (r *QueryResults) Concatenate(itemList *QueryResults) {
	*r = append(*r, *itemList...)
}

func (r *QueryResults) First() ComparableResultItem {
	return (*r)[0]
}

func (r *QueryResults) Length() int {
	return len(*r)
}

func (r *QueryResults) RemoveDuplicates() {
	keys := make(map[string]bool)
	list := EmptyQueryResult()
	for _, entry := range *r {
		if _, value := keys[entry.GetCompareField()]; !value {
			keys[entry.GetCompareField()] = true
			list.Add(entry)
		}
	}
	*r = *list
}

func (r *QueryResults) DuplicateElements() *QueryResults {
	duplicateElements := EmptyQueryResult()

	keys := make(map[string]bool)
	for _, entry := range *r {
		if _, value := keys[entry.GetCompareField()]; !value {
			keys[entry.GetCompareField()] = true
		} else {
			duplicateElements.Add(entry)
		}
	}
	duplicateElements.RemoveDuplicates()
	return duplicateElements
}

func (r *QueryResults) Limit(limitSize int) *QueryResults {
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
