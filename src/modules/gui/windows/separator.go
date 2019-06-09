package windows

import . "github.com/lxn/walk/declarative"

type MySeparator struct {
	WalkSeparator Separator
}

func NewSeparator() *MySeparator {
	return &MySeparator{}
}

func (m *MySeparator) Create() Separator {
	return m.WalkSeparator
}
