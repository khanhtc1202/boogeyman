package domain

import "strings"

type Keyword string

func NewKeyword(value string) Keyword {
	return Keyword(value)
}

func (k Keyword) String() string {
	return strings.Replace(string(k), " ", "%20", -1)
}
