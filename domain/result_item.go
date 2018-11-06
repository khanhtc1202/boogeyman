package domain

type ComparableResultItem interface {
	GetCompareField() string
}

// TODO make interface
type UrlBaseResultItem struct {
	ComparableResultItem
	createdTime string
	title       string
	description string
	url         string
}

func NewResultItem(
	createdTime string,
	title string,
	description string,
	url string,
) *UrlBaseResultItem {
	return &UrlBaseResultItem{
		createdTime: createdTime,
		title:       title,
		description: description,
		url:         url,
	}
}

func (r *UrlBaseResultItem) Time() string {
	return r.createdTime
}

func (r *UrlBaseResultItem) GetTitleString() string {
	return r.title
}

func (r *UrlBaseResultItem) GetDescription() string {
	return r.description
}

func (r *UrlBaseResultItem) GetUrl() string {
	return r.url
}

func (r *UrlBaseResultItem) GetCompareField() string {
	return r.url
}
