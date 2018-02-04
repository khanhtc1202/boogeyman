package main

import (
	"fmt"

	"os"

	"github.com/khanhtc1202/boogeyman/adapter/controller"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/service"
	"github.com/khanhtc1202/boogeyman/domain"
	spiderPool "github.com/khanhtc1202/boogeyman/infrastructure/service"
)

func main() {
	bingSpider := spiderPool.NewBingSpider()
	googleSpider := spiderPool.NewGoogleSpider()
	materialPool := repository.NewMaterialPool([]service.Collector{googleSpider, bingSpider})

	params := parseCommandParams()
	keyword := domain.NewKeyword(params[0])
	materialPool.Fetch(keyword)

	//materialPool.Fetch(domain.NewKeyword("java"))

	boogeyman := controller.NewBoogeyman(materialPool)

	results, err := boogeyman.ShowSearchResult(domain.ALL, materialPool.GetSearchEngineList())
	if err != nil {
		fmt.Println("Error : ", err)
		panic("Error on show search results!")
	}
	for _, result := range *results {
		fmt.Println(result.Show())
	}
}

func parseCommandParams() []string {
	if len(os.Args) < 2 {
		panic("Missing params")
	}
	return os.Args[1:]
}
