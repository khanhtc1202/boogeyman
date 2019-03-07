package presenter

import "github.com/khanhtc1202/boogeyman/internal/domain"

type TextPresenter interface {
	PrintList(results *domain.QueryResult) error
}
