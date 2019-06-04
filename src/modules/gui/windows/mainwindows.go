package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
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

func (m *MyMainWindow) Font(font Font)*MyMainWindow {
	m.mw.Font = font
	return m
}

func