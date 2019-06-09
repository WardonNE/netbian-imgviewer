package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"fmt"
)

type MyMainWindow struct {
	*walk.MainWindow
	mw MainWindow
}

func NewMainWindow(mw MainWindow) *MyMainWindow {
	return &MyMainWindow{mw: mw}
}

func (m *MyMainWindow) AssignTo() *MyMainWindow {
	m.mw.AssignTo = &m.MainWindow
	return m
}

func (m *MyMainWindow) Title(title string) *MyMainWindow {
	m.mw.Title = title
	return m
}

func (m *MyMainWindow) Icon(icon string) *MyMainWindow {
	m.mw.Icon = icon
	return m
}

func (m *MyMainWindow) Name(name string) *MyMainWindow {
	m.mw.Name = name
	return m
}

func (m *MyMainWindow) Enabled(enabled bool) *MyMainWindow {
	m.mw.Enabled = enabled
	return m
}

func (m *MyMainWindow) Font(font Font) *MyMainWindow {
	m.mw.Font = font
	return m
}

func (m *MyMainWindow) Size(size Size) *MyMainWindow {
	m.mw.Size = size
	return m
}

func (m *MyMainWindow) MinSize(s Size) *MyMainWindow {
	m.mw.MinSize = s
	return m
}

func (m *MyMainWindow) MaxSize(s Size) *MyMainWindow {
	m.mw.MaxSize = s
	return m
}

func (m *MyMainWindow) VBoxLayout(vbox VBox) *MyMainWindow {
	m.mw.Layout = vbox
	return m
}

func (m *MyMainWindow) HBosxLayout(hbox HBox) *MyMainWindow {
	m.mw.Layout = hbox
	return m
}

func (m *MyMainWindow) Children(w []Widget) *MyMainWindow {
	m.mw.Children = w
	return m
}

func (m *MyMainWindow) MenuItems(menulist []MenuItem) *MyMainWindow {
	m.mw.MenuItems = menulist
	return m
}

func (m *MyMainWindow) Run() {
	if _, err := m.mw.Run(); err != nil {
		fmt.Printf("Run Main Window Error: \r\n %v", err)
	}
}
