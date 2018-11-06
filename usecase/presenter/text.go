package presenter

import "github.com/khanhtc1202/boogeyman/domain"

type TextPresenter interface {
	PrintList(results *domain.QueryResult) error
}
