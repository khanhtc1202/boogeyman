package main

import (
	"fmt"

	"github.com/khanhtc1202/boogeyman/adapter/controller"
	"github.com/khanhtc1202/boogeyman/adapter/persistent/repository"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/domain/search_engine"
	"github.com/khanhtc1202/boogeyman/infrastructure/service"
)

func main() {
	searchEngineList := search_engine.EmptySearchEngineList()
	searchEngineList.Add(search_engine.BING)

	keyword := domain.NewKeyword("sample")
	bingSpider := service.NewBingSpider()
	materialPool := repository.NewMaterialPool(bingSpider)
	materialPool.Fetch(keyword, searchEngineList)

	boogeyman := controller.NewBoogeyman(materialPool, searchEngineList)

	results, err := boogeyman.ShowSearchResult(domain.TOP)
	if err != nil {
		fmt.Println("Error : ", err)
		panic("Error on show search results!")
	}
	for _, result := range *results {
		fmt.Println(result.Show())
	}
}
