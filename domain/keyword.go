package domain

type Keyword struct {
	value string
}

func NewKeyword(keyword string) *Keyword {
	return &Keyword{
		value: keyword,
	}
}

func (k *Keyword) String() string {
	return k.value
}
