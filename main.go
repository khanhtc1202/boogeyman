package main

import (
	"fmt"

	"github.com/khanhtc1202/boogeyman/adapter/controller"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	spiderPool "github.com/khanhtc1202/boogeyman/infrastructure/service"
)

func main() {
	keyword := domain.NewKeyword("sample")
	bingSpider := spiderPool.NewBingSpider()
	googleSpider := spiderPool.NewGoogleSpider()

	materialPool := repository.NewMaterialPool([]service.Collector{bingSpider, googleSpider})
	materialPool.Fetch(keyword)

	boogeyman := controller.NewBoogeyman(materialPool)

	results, err := boogeyman.ShowSearchResult(domain.TOP, materialPool.GetSearchEngineList())
	if err != nil {
		fmt.Println("Error : ", err)
		panic("Error on show search results!")
	}
	for _, result := range *results {
		fmt.Println(result.Show())
	}
}
