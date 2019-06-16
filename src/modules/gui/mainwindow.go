package gui

import (
	"modules/app"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow

	catalogcb *walk.ComboBox
}

func CreateMainWindow() {
	var mw = &MyMainWindow{}
	m := MainWindow{
		AssignTo:  &mw.MainWindow,
		Title:     getTitle(),
		Size:      getSize(),
		MinSize:   getMinSize(),
		MaxSize:   getMaxSize(),
		MenuItems: GetMenuList(),
		Layout:    GetVBox(),
		Children:  GetMainWindowChildren(),
	}
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
