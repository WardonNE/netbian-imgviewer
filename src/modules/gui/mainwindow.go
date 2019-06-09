package gui

import (
	"modules/app"
	"modules/gui/windows"

	. "github.com/lxn/walk/declarative"
)

var mw MainWindow

func CreateMainWindow() {
	m := windows.NewMainWindow(mw).Title(getTitle())
	m.Size(getSize()).MinSize(getMinSize()).MaxSize(getMaxSize())
	m.MenuItems(GetMenuList())
	m.VBoxLayout(GetVBox())
	m.Run()
}

func getTitle() string {
	return app.MainWindowConf.Profiles.Title + " " + app.MainWindowConf.Profiles.Version + " -- " + app.MainWindowConf.Profiles.Author
}

func getSize() Size {
	return Size{
		app.MainWindowConf.MainWindowSize.Size.Width,
		app.MainWindowConf.MainWindowSize.Size.Height,
	}
}

func getMinSize() Size {
	return Size{
		app.MainWindowConf.MainWindowSize.MinSize.Width,
		app.MainWindowConf.MainWindowSize.MinSize.Height,
	}
}

func getMaxSize() Size {
	return Size{
		app.MainWindowConf.MainWindowSize.MaxSize.Width,
		app.MainWindowConf.MainWindowSize.MaxSize.Height,
	}
}
