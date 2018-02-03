package domain

type ResultItems []ResultItem

type ResultItem struct {
	createdTime string
	title       string
	description string
	url         string
}

func (r *ResultItem) Time() string {
	return r.createdTime
}

func (r *ResultItem) GetTitleString() string {
	return r.title
}

func (r *ResultItem) GetDescription() string {
	return r.description
}

func (r *ResultItem) GetUrl() string {
	return r.url
}

func (r *ResultItems) Add(resultItem *ResultItem) {
	if resultItem == nil {
		return
	}
	*r = append(*r, *resultItem)
}
