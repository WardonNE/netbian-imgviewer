package windows

import (
	. "github.com/lxn/walk/declarative"
)

type MyMenu struct {
	WalkMenu Menu
}

func NewMenu() *MyMenu {
	return &MyMenu{}
}

func (m *MyMenu) SetText(t string) *MyMemu {
	m.WalkMenu.Text = t
	return m
}

func (m *MyMenu) SetImage(img interface{}) *MyMenu {
	m.WalkMenu.Image = img
	return m
}

func (m *MyMenu) SetEnabled(e Property) *MyMenu {
	m.WalkMenu.Enabled = e
	return m
}

func (m *MyMenu) SetVisible(v Property) *MyMenu {
	m.WalkMenu.Visible = v
	return m
}

func (m *MyMenu) SetMenuItem(items []MenuItem) *MyMenu {
	m.WalkMenu.Items = items
	return m
}

func (m *MyMenu) Create() Menu {
	return m.WalkMenu
}
