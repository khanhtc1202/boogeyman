package domain

type ResultItems []*ResultItem

func EmptyResultItems() *ResultItems {
	return &ResultItems{}
}

func (r *ResultItems) Add(resultItem *ResultItem) {
	if resultItem == nil {
		return
	}
	*r = append(*r, resultItem)
}

func (r *ResultItems) Concatenate(itemList *ResultItems) {
	for _, item := range *itemList {
		r.Add(item)
	}
}

func (r *ResultItems) First() *ResultItem {
	return (*r)[0]
}

func (r *ResultItems) RemoveDuplicates() {
	keys := make(map[string]bool)
	list := EmptyResultItems()
	for _, entry := range *r {
		if _, value := keys[entry.url]; !value {
			keys[entry.url] = true
			list.Add(entry)
		}
	}
	*r = *list
}

func (r *ResultItems) DuplicateElements() *ResultItems {
	duplicateElements := EmptyResultItems()

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

func (r *ResultItems) Limit(limitSize int) *ResultItems {
	splitSlide := EmptyResultItems()

	if limitSize > len(*r) {
		return r
	} else {
		for _, item := range (*r)[0:limitSize] {
			splitSlide.Add(item)
		}
		return splitSlide
	}
}
