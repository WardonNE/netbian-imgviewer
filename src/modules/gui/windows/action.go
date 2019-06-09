package windows

import (
	. "github.com/lxn/walk/declarative"
)

type MyAction struct {
	WalkAction Action
}

func NewAction() *MyAction {
	return &MyAction{}
}

func (m *MyAction) Text(text string) *MyAction {
	m.WalkAction.Text = text
	return m
}

func (m *MyAction) Image(image interface{}) *MyAction {
	m.WalkAction.Image = image
	return m
}

func (m *MyAction) Checked(checked bool) *MyAction {
	m.WalkAction.Checked = checked
	return m
}

func (m *MyAction) Enabled(enabled bool) *MyAction {
	m.WalkAction.Enabled = enabled
	return m
}

func (m *MyAction) Visible(visible bool) *MyAction {
	m.WalkAction.Visible = visible
	return m
}

func (m *MyAction) Shortcut(shortcut Shortcut) *MyAction {
	m.WalkAction.Shortcut = shortcut
	return m
}

func (m *MyAction) Checkable(checkable bool) *MyAction {
	m.WalkAction.Checkable = checkable
	return m
}

func (m *MyAction) Create() Action {
	return m.WalkAction
}
