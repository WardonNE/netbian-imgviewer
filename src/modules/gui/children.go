package gui

import (
	. "github.com/lxn/walk/declarative"
)

func GetMainWindowChildren() []Widget {
	return []Widget{
		GetToolBarList(),
		GetBodyComposite(),
	}
}
