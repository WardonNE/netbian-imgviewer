package window

import (
	"modules/app"
	"modules/util/gui/windows"

	. "github.com/lxn/walk/declarative"
)

var (
	mw   MainWindow
	conf *app.MainWindowConfig
)

func init() {
	loadConfig()
}

func Run() {
	mmw := windows.NewMainWindow(mw).AssignTo()
	mmw = mmw.SetTitle(conf.Title)
	mmw = mmw.SetSize(windows.NewSize().SetWidth(conf.Width).SetHeight(conf.Height).Create())
	mmw = mmw.SetMinSize(windows.NewSize().SetWidth(conf.MinWidth).SetHeight(conf.MinHeight).Create())
	mmw = mmw.SetMaxSize(windows.NewSize().SetWidth(conf.MaxWidth).SetHeight(conf.MaxHeight).Create())
	mmw.Run()
}

func loadConfig() {
	conf = app.MainWindowConfMgr.ConfigValues.Load().(*app.MainWindowConfig)
}
