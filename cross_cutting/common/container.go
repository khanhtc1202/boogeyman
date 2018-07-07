package common

import (
	"strings"

	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/adapter/presenter/console"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/infrastructure/io/data"
	spiderPool "github.com/khanhtc1202/boogeyman/infrastructure/service"
	"github.com/khanhtc1202/boogeyman/usecase/interactor"
)

type DIContainer struct {
	cmdParams *data.CommandParams
}

func NewDIContainer(
	cmdParams *data.CommandParams,
) *DIContainer {
	return &DIContainer{
		cmdParams: cmdParams,
	}
}

func (d *DIContainer) SearchInfo() *interactor.InfoSearch {
	textPresenter := console.NewColorfulTextPresenter()
	materialPoolRepo := materialPoolFactory(d.cmdParams.Engine)

	return interactor.NewInfoSearch(textPresenter, materialPoolRepo)
}

func (d *DIContainer) GetQueryStrategy() domain.RankerStrategyType {
	switch strings.ToUpper(d.cmdParams.Strategy) {
	case domain.TOP.String():
		return domain.TOP
	case domain.CROSS.String():
		return domain.CROSS
	default:
		return domain.ALL
	}
}

func materialPoolFactory(selectedEngine string) *repository.QueryResultPool {
	collectors := service.EmptyCollectorList()
	switch strings.ToUpper(selectedEngine) {
	case domain.GOOGLE.String():
		collectors.Add(spiderPool.NewGoogleSpider())
		break
	case domain.BING.String():
		collectors.Add(spiderPool.NewBingSpider())
		break
	case domain.ASK.String():
		collectors.Add(spiderPool.NewAskSpider())
		break
	default:
		collectors.Add(spiderPool.NewAskSpider())
		collectors.Add(spiderPool.NewBingSpider())
		collectors.Add(spiderPool.NewGoogleSpider())
	}
	return repository.NewResultPool(*collectors)
}
