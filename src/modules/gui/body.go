package gui

import (
	"fmt"
	"log"
	"modules/app"
	"modules/crawl"
	"strings"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func GetBodyComposite() Composite {
	return Composite{
		Layout: GetBodyGrid(),
		Children: []Widget{
			GetImageListBox(),
			GetImageView(),
		},
	}
}

func GetBodyGrid() Grid {
	grid := app.MainWindowConf.Children.BodyComposite.Layout.Grid
	return Grid{
		Columns: grid.Columns,
		Spacing: grid.Spacing,
	}
}

func GetImageListBox() ListBox {
	imagelb := app.MainWindowConf.Children.BodyComposite.ImageListBox
	return ListBox{
		AssignTo: &mw.imagelb,
		MaxSize: Size{
			imagelb.MaxSize.Width,
			imagelb.MaxSize.Height,
		},
		MinSize: Size{
			imagelb.MinSize.Width,
			imagelb.MinSize.Height,
		},
		Font: Font{
			Family:    imagelb.Font.Family,
			PointSize: imagelb.Font.PointSize,
			Bold:      imagelb.Font.Bold,
			StrikeOut: imagelb.Font.StrikeOut,
			Italic:    imagelb.Font.Italic,
			Underline: imagelb.Font.UnderLine,
		},
		Row:                   imagelb.Row,
		RowSpan:               imagelb.RowSpan,
		Column:                imagelb.Column,
		ColumnSpan:            imagelb.ColumnSpan,
		AlwaysConsumeSpace:    imagelb.AlwaysConsumeSpace,
		Background:            SolidColorBrush{Color: walk.RGB(imagelb.Background.R, imagelb.Background.G, imagelb.Background.B)},
		Model:                 mw.imagelbmodel,
		OnCurrentIndexChanged: mw.imageLbModelOnCurrentIndexChanged,
	}
}

type imageListBoxModel struct {
	walk.ListModelBase
	items []imageItem
}

type imageItem struct {
	title, url string
}

func GetImageListBoxModel(page int) *imageListBoxModel {
	imagelist, activepage, totalpage := crawl.LoadImageList(page)
	model := &imageListBoxModel{
		items: make([]imageItem, len(imagelist)),
	}
	for key, image := range imagelist {
		model.items[key] = imageItem{
			title: image.Name,
			url:   image.Url,
		}
	}
	mw.imagelbmodel = model
	mw.activepage = activepage
	mw.totalpage = totalpage
	return model
}

func reloadImageList(page int) {
	imagelist, activepage, totalpage := crawl.LoadImageList(page)
	model := &imageListBoxModel{
		items: make([]imageItem, len(imagelist)),
	}
	for key, image := range imagelist {
		model.items[key] = imageItem{
			title: image.Name,
			url:   image.Url,
		}
	}
	mw.imagelbmodel = model
	mw.activepage = activepage
	mw.totalpage = totalpage
	mw.imagelb.SetModel(model)
}

func reloadImageListModelByCatalog(url string, page int) {
	imagelist, activepage, totalpage := crawl.LoadImageByCatalog(url, page)
	model := &imageListBoxModel{
		items: make([]imageItem, len(imagelist)),
	}
	for key, image := range imagelist {
		model.items[key] = imageItem{
			title: image.Name,
			url:   image.Url,
		}
	}
	mw.imagelbmodel = model
	mw.imagelb.SetModel(model)
	mw.sizecb.SetCurrentIndex(-1)
	mw.activepage = activepage
	mw.totalpage = totalpage
}

func reloadImageListModelBySize(url string, page int) {
	imagelist, activepage, totalpage := crawl.LoadImageBySize(url, page)
	model := &imageListBoxModel{
		items: make([]imageItem, len(imagelist)),
	}
	for key, image := range imagelist {
		model.items[key] = imageItem{
			title: image.Name,
			url:   image.Url,
		}
	}
	mw.imagelbmodel = model
	mw.imagelb.SetModel(model)
	mw.catalogcb.SetCurrentIndex(-1)
	mw.activepage = activepage
	mw.totalpage = totalpage
}

func reloadIamgeListModeBySearchKeyword(keyword string, page int) {
	imagelist, activepage, totalpage := crawl.LoadImageBySearchKeyword(keyword, page)
	fmt.Println(activepage, totalpage)
	model := &imageListBoxModel{
		items: make([]imageItem, len(imagelist)),
	}
	for key, image := range imagelist {
		model.items[key] = imageItem{
			title: image.Name,
			url:   image.Url,
		}
	}
	mw.imagelbmodel = model
	mw.imagelb.SetModel(model)
	mw.catalogcb.SetCurrentIndex(-1)
	mw.sizecb.SetCurrentIndex(-1)
	mw.activepage = activepage
	mw.totalpage = totalpage
}

func (i *imageListBoxModel) ItemCount() int {
	return len(i.items)
}

func (i *imageListBoxModel) Value(index int) interface{} {
	return i.items[index].title
}

func (mw *MyMainWindow) imageLbModelOnCurrentIndexChanged() {
	i := mw.imagelb.CurrentIndex()
	fmt.Println("Current Index(Image List Box): ", i)
	if i < 0 {
		return
	}
	activeItem := mw.imagelbmodel.items[i]
	filename := crawl.DownloadImage(activeItem.url, activeItem.title)
	if filename != "" {
		filename = strings.Replace(filename, "/", "\\", -1)
		filename = strings.Replace(filename, "bin\\..\\", "", 1)
		fmt.Println(filename)
		image, err := walk.NewImageFromFile(filename)
		if err != nil {
			log.Panicln("new image error: ", err)
		}
		err = mw.imageviewer.SetImage(image)
		if err != nil {
			log.Panicln("set image error: ", err)
		}
	}
}
