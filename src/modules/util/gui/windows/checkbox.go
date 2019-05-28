package windows

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyCheckBox struct {
	WalkCheckBox CheckBox
}

func NewCheckBox() *MyCheckBox {
	return &MyCheckBox{}
}

func (m *MyCheckBox) SetBackGround(bg Brush) *MyCheckBox {
	m.WalkCheckBox.Background = bg
	return m
}

func (m *MyCheckBox) SetContextMenuItems(mi []MenuItem) *MyCheckBox {
	m.WalkCheckBox.ContextMenuItems = mi
	return m
}

func (m *MyCheckBox) SetEnabeld(e bool) *MyCheckBox {
	m.WalkCheckBox.Enabled = e
	return m
}

func (m *MyCheckBox) SetFont(f Font) *MyCheckBox {
	m.WalkCheckBox.Font = f
	return m
}

func (m *MyCheckBox) SetMaxSize(s Size) *MyCheckBox {
	m.WalkCheckBox.MaxSize = s
	return m
}

func (m *MyCheckBox) SetMinSize(s Size) *MyCheckBox {
	m.WalkCheckBox.MinSize = s
	return m
}

func (m *MyCheckBox) SetName(n string) *MyCheckBox {
	m.WalkCheckBox.Name = n
	return m
}

func (m *MyCheckBox) SetPresistent(p bool) *MyCheckBox {
	m.WalkCheckBox.Persistent = p
	return m
}

func (m *MyCheckBox) SetRightToLeftReading(r bool) *MyCheckBox {
	m.WalkCheckBox.RightToLeftReading = r
	return m
}

func (m *MyCheckBox) SetToolTipText(t Property) *MyCheckBox {
	m.WalkCheckBox.ToolTipText = t
	return m
}

func (m *MyCheckBox) SetVisible(v Property) *MyCheckBox {
	m.WalkCheckBox.Visible = v
	return m
}

func (m *MyCheckBox) SetAlwaysConsumeSpace(a bool) *MyCheckBox {
	m.WalkCheckBox.AlwaysConsumeSpace = a
	return m
}

func (m *MyCheckBox) SetColumn(c int) *MyCheckBox {
	m.WalkCheckBox.Column = c
	return m
}

func (m *MyCheckBox) SetColumnSpan(c int) *MyCheckBox {
	m.WalkCheckBox.ColumnSpan = c
	return m
}

func (m *MyCheckBox) SetGraphicsEffects(g []walk.WidgetGraphicsEffect) *MyCheckBox {
	m.WalkCheckBox.GraphicsEffects = g
	return m
}

func (m *MyCheckBox) SetRow(r int) *MyCheckBox {
	m.WalkCheckBox.Row = r
	return m
}

func (m *MyCheckBox) SetRowSpan(r int) *MyCheckBox {
	m.WalkCheckBox.RowSpan = r
	return m
}

func (m *MyCheckBox) SetStretchFactor(s int) *MyCheckBox {
	m.WalkCheckBox.StretchFactor = s
	return m
}

func (m *MyCheckBox) SeteChecked(c Property) *MyCheckBox {
	m.WalkCheckBox.Checked = c
	return m
}

func (m *MyCheckBox) SetText(t Property) *MyCheckBox {
	m.WalkCheckBox.Text = t
	return m
}

func (m *MyCheckBox) SetCheckState(c Property) *MyCheckBox {
	m.WalkCheckBox.CheckState = c
	return m
}

func (m *MyCheckBox) SetTextOnLeftSide(t bool) *MyCheckBox {
	m.WalkCheckBox.TextOnLeftSide = t
	return m
}

func (m *MyCheckBox) Tristate(t bool) *MyCheckBox {
	m.WalkCheckBox.Tristate = t
	return m
}

func (m *MyCheckBox) Create() CheckBox {
	return m.WalkCheckBox
}
