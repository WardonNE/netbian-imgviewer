package gui

import (
	"modules/app"

	. "github.com/lxn/walk/declarative"
)

func GetMenuList() []MenuItem {
	return []MenuItem{GetAboutMenu()}
}

func GetAboutMenu() Menu {
	aboutmenu := app.MainWindowConf.MenuList.AboutMenu
	return Menu{
		Text:  aboutmenu.Text,
		Items: GetAboutMenuList(),
	}
}

func GetAboutMenuList() []MenuItem {
	return []MenuItem{
		GetAboutMenuSeparator(),
		AboutViewAction(),
	}
}

func GetAboutMenuSeparator() Separator {
	//separator := app.MainWindowConf.MenuList.AboutMenu.Items.MenuList.MenuList.Separator
	return Separator{}
}

func AboutViewAction() Action {
	action := app.MainWindowConf.MenuList.AboutMenu.Items.MenuList.Action
	return Action{
		Text: action.Text,
	}
}
