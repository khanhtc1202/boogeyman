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

func (r *ResultItems) First() *ResultItem {
	return (*r)[0]
}
