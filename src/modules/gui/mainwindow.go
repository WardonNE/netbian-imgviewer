package gui

import (
	"log"
	"modules/app"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	activepage            int
	totalpage             int
	searchkeyword         string
	catalogcbcurrentindex int
	sizecbcurrentindex    int
	catalogcb             *walk.ComboBox
	catalogcbmodel        *catalogComboBoxModel
	sizecb                *walk.ComboBox
	sizecbmodel           *sizeComboBoxModel
	imagelb               *walk.ListBox
	imagelbmodel          *imageListBoxModel
	searchpb              *walk.PushButton
	searchle              *walk.LineEdit
	nextpagepb            *walk.PushButton
	prevpagepb            *walk.PushButton
}

var mw = &MyMainWindow{
	activepage: 1,
}

func init() {
	GetImageListBoxModel(1)
}

func CreateMainWindow() {
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
	if _, err := m.Run(); err != nil {
		log.Panicln("run mainwindow error:", err)
	}
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
