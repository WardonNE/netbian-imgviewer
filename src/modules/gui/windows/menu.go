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

func (m *MyMenu) Text(t string) *MyMenu {
	m.WalkMenu.Text = t
	return m
}

func (m *MyMenu) Image(img interface{}) *MyMenu {
	m.WalkMenu.Image = img
	return m
}

func (m *MyMenu) Enabled(e Property) *MyMenu {
	m.WalkMenu.Enabled = e
	return m
}

func (m *MyMenu) Visible(v Property) *MyMenu {
	m.WalkMenu.Visible = v
	return m
}

func (m *MyMenu) MenuItem(items []MenuItem) *MyMenu {
	m.WalkMenu.Items = items
	return m
}

func (m *MyMenu) Create() Menu {
	return m.WalkMenu
}
