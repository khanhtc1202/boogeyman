package domain

import "strings"

type Keyword struct {
	value string
}

func NewKeyword(keyword string) *Keyword {
	return &Keyword{
		value: keyword,
	}
}

func (k *Keyword) String() string {
	return strings.Replace(k.value, " ", "%20", -1)
}
