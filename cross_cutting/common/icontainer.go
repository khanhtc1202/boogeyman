package common

import "github.com/khanhtc1202/boogeyman/usecase/interactor"

type IDIContainer interface {
	SearchInfo() *interactor.InfoSearch
}
