package console

import (
	"github.com/fatih/color"
	ioInterface "github.com/khanhtc1202/boogeyman/cross_cutting/io"
	"github.com/khanhtc1202/boogeyman/domain"
	"github.com/khanhtc1202/boogeyman/infrastructure/io"
)

type TextPresenter struct {
	writer ioInterface.UI
}

func NewColorfulTextPresenter() *TextPresenter {
	return &TextPresenter{
		writer: io.ColorfulConsole(),
	}
}

func (t *TextPresenter) PrintList(results *domain.QueryResult) {
	for _, result := range *results {
		t.presentItem(result)
	}
	t.writer.Printf(color.HiCyanString("\nTotal %v result(s) founded!\n", len(*results)))
}

func (t *TextPresenter) presentItem(result *domain.ResultItem) {
	t.writer.Printf(color.HiGreenString("Title: %v \n", result.GetTitleString()))
	t.writer.Printf(color.YellowString("URL: %v \n", result.GetUrl()))
	t.writer.Printf(color.RedString("Description: ") + result.GetDescription() + "\n")
	t.writer.Printf(color.BlueString("Create At: %v \n", result.Time()))
	t.writer.Println("---------------------")
}
