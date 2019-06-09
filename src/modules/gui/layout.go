package gui

import (
	. "github.com/lxn/walk/declarative"

	"modules/app"
)

func GetVBox() VBox {
	return VBox{
		MarginsZero: app.MainWindowConf.VBox.MarginsZero,
	}
}
