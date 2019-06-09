package gui

import (
	"modules/app"
	"modules/gui/windows"

	. "github.com/lxn/walk/declarative"
)

func GetMenuList() []MenuItem {
	return []MenuItem{GetAboutMenu()}
}

func GetAboutMenu() Menu {
	m := windows.NewMenu()
	m.Text(app.MainWindowConf.MenuList.AboutMenu.Text)
	m.MenuItem(GetAboutMenuList())
	return m.Create()
}

func GetAboutMenuList() []MenuItem {
	return []MenuItem{
		GetAboutMenuSeparator(),
		AboutViewAction(),
	}
}

func GetAboutMenuSeparator() Separator {
	return windows.NewSeparator().Create()
}

func AboutViewAction() Action {
	m := windows.NewAction()
	m.Text(app.MainWindowConf.MenuList.AboutMenu.Items.MenuList.Action.Text)
	return m.Create()
}
