package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	// "fmt"
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

func (m *MyMainWindow) SetTitle(t string) *MyMainWindow {
	m.mw.Title = t
	return m
}

func (m *MyMainWindow) SetIcon(i string) *MyMainWindow {
	m.mw.Icon = i
	return m
}

func (m *MyMainWindow) SetName(n string) *MyMainWindow {
	m.mw.Name = n
	return m
}

func (m *MyMainWindow) SetEnabled(e bool) *MyMainWindow {
	m.mw.Enabled = e
	return m
}

func (m *MyMainWindow) SetFont(f Font) *MyMainWindow {
	m.mw.Font = f
	return m
}

func (m *MyMainWindow) SetSize(s Size) *MyMainWindow {
	m.mw.Size = s
	return m
}

func (m *MyMainWindow) SetMinSize(s Size) *MyMainWindow {
	m.mw.MinSize = s
	return m
}

func (m *MyMainWindow) SetMaxSize(s Size) *MyMainWindow {
	m.mw.MaxSize = s
	return m
}

func (m *MyMainWindow) SetVBoxLayout(vbox VBox) *MyMainWindow {
	m.mw.Layout = vbox
	return m
}

func (m *MyMainWindow) SetHBosxLayout(hbox HBox) *MyMainWindow {
	m.mw.Layout = hbox
	return m
}

func (m *MyMainWindow) SetChildren(w []Widget) *MyMainWindow {
	m.mw.Children = w
	return m
}

func (m *MyMainWindow) Run() {
	if _, err := m.mw.Run(); err != nil {
		panic(err)
	}
}
