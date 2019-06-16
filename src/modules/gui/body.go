package gui

import (
	"modules/app"

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
		MaxSize: Size{
			imagelb.MaxSize.Width,
			imagelb.MaxSize.Height,
		},
		MinSize: Size{
			imagelb.MinSize.Width,
			imagelb.MinSize.Height,
		},
	}
}
