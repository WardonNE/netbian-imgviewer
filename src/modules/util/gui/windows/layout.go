package windows

import (
	. "github.com/lxn/walk/declarative"
)

type MySize struct {
	WalkSize Size
}

func NewSize() *MySize {
	return &MySize{}
}

func (m *MySize) SetWidth(w int) *MySize {
	m.WalkSize.Width = w
	return m
}

func (m *MySize) SetHeight(h int) *MySize {
	m.WalkSize.Height = h
	return m
}

func (m *MySize) Create() Size {
	return m.WalkSize
}

type MyMargins struct {
	WalkMargins Margins
}

func NewMargins() *MyMargins {
	return &MyMargins{}
}

func (m *MyMargins) SetLeft(l int) *MyMargins {
	m.WalkMargins.Left = l
	return m
}

func (m *MyMargins) SetRight(r int) *MyMargins {
	m.WalkMargins.Right = r
	return m
}

func (m *MyMargins) SetTop(t int) *MyMargins {
	m.WalkMargins.Top = t
	return m
}

func (m *MyMargins) SetBottom(b int) *MyMargins {
	m.WalkMargins.Bottom = b
	return m
}

func (m *MyMargins) Create() *MyMargins {
	return m.WalkMargins
}

type MyHBox struct {
	WalkHBox HBox
}

func NewHBox() *MyHBox {
	return &MyHBox{}
}

func (m *MyHBox) SetMargins(margins Margins) *MyHBox {
	m.WalkHBox.Margins = margins
	return m
}

func (m *MyHBox) SetSpacing(s int) *MyHBox {
	m.WalkHBox.Spacing = s
	return m
}

func (m *MyHBox) SetMarginsZero(z bool) *MyHBox {
	m.WalkHBox.MarginsZero = z
	return m
}

func (m *MyHBox) SetSpacingZere(s bool) *MyHBox {
	m.WalkHBox.SpacingZero = s
	return m
}

func (m *MyHBox) Create() HBox {
	return m.WalkHBox
}

type MyVBox struct {
	WalkVBox VBox
}

func NewVBox() *MyVBox {
	return &MyVBox{}
}

func (m *MyVBox) SetMargins(margins Margins) *MyVBox {
	m.WalkVBox.Margins = margins
	return m
}

func (m *MyVBox) SetSpacing(s int) *MyVBox {
	m.WalkVBox.Spacing = s
	return m
}

func (m *MyVBox) SetMarginsZero(z bool) *MyVBox {
	m.WalkVBox.MarginsZero = z
	return m
}

func (m *MyVBox) SetSpacingZere(s bool) *MyVBox {
	m.WalkVBox.SpacingZero = s
	return m
}

func (m *MyVBox) Create() VBox {
	return m.WalkVBox
}
