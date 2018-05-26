package domain

type ResultItem struct {
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
) *ResultItem {
	return &ResultItem{
		createdTime: createdTime,
		title:       title,
		description: description,
		url:         url,
	}
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

func (r *ResultItem) Show() string {
	buildString := "\n =>> Title: " + r.GetTitleString() +
		" \n =>> URL: " + r.GetUrl() +
		" \n =>> Description: " + r.GetDescription() +
		" \n =>> Create At: " + r.Time()
	return buildString
}
