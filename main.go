package main

import (
	"os"

	"github.com/khanhtc1202/boogeyman/adapter/controller"
	"github.com/khanhtc1202/boogeyman/cross_cutting/common"
	"github.com/khanhtc1202/boogeyman/cross_cutting/io"
	io2 "github.com/khanhtc1202/boogeyman/infrastructure/io"
	"github.com/khanhtc1202/boogeyman/infrastructure/meta_info"
)

var (
	version   string
	revision  string
	buildDate string
	goVersion string
	mode      string
)

var metaInfo = meta_info.NewMetaInfo(
	version,
	revision,
	buildDate,
	goVersion,
	mode,
)

func main() {
	commandParse := io2.NewCommandParse()

	// parse command params
	cmdParams := commandParse.ParseCommandParams()

	// check meta_info
	if cmdParams.ShowVersion {
		ShowMetaInfo(metaInfo)
	}

	container := common.NewDIContainer(cmdParams)
	queryStrategy := container.GetQueryStrategy()

	boogeyman := controller.NewBoogeyman(container)

	err := boogeyman.Search(cmdParams.QueryString, queryStrategy)
	if err != nil {
		io.Errorln(err)
		os.Exit(1)
	}
}

func ShowMetaInfo(metaInfo *meta_info.MetaInfo) {
	io.Infof(metaInfo.GetMetaInfo())
	os.Exit(0)
}
