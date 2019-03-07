package console

import (
	"github.com/fatih/color"
	ioInterface "github.com/khanhtc1202/boogeyman/internal/cross_cutting/io"
	"github.com/khanhtc1202/boogeyman/internal/domain"
	"github.com/khanhtc1202/boogeyman/internal/infrastructure/io"
	"github.com/pkg/errors"
)

type TextPresenter struct {
	writer ioInterface.UI
}

func NewColorfulTextPresenter() *TextPresenter {
	return &TextPresenter{
		writer: io.ColorfulConsole(),
	}
}

func (t *TextPresenter) PrintList(results *domain.QueryResult) error {
	for _, result := range *results {
		switch result.(type) {
		case *domain.UrlBaseResultItem:
			t.presentUrlBaseItem(result.(*domain.UrlBaseResultItem))
			continue
		default:
			return errors.New("Error not found presenter for this type of ResultItem")
		}
	}

	t.writer.Printf(color.HiCyanString("\nTotal %v result(s) founded!\n", len(*results)))
	return nil
}

func (t *TextPresenter) presentUrlBaseItem(result *domain.UrlBaseResultItem) {
	t.writer.Printf(color.HiGreenString("Title: %v \n", result.GetTitleString()))
	t.writer.Printf(color.YellowString("URL: %v \n", result.GetUrl()))
	t.writer.Printf(color.RedString("Description: ") + result.GetDescription() + "\n")
	t.writer.Printf(color.BlueString("Create At: %v \n", result.Time()))
	t.writer.Println("---------------------")
}
