package windows

import . "github.com/lxn/walk/declarative"

type MyPushButton struct {
	WalkPushButton PushButton
}

func NewPushButton() *MyPushButton {
	return &MyPushButton{}
}

func (m *MyPushButton) ContextMenuItems(menuitems []MenuItem) {
	m.WalkPushButton.ContextMenuItems = menuitems
	return m
}

func (m *MyPushButton) DoubleBuffering(dbuffering bool) *MyPushButton {
	m.WalkPushButton.DoubleBuffering = dbuffering
	return m
}

func (m *MyPushButton) Enabled(enabled bool) *MyPushButton {
	m.WalkPushButton.Enabled = enabled
	return m
}

func (m *MyPushButton) Font(font Font) *MyPushButton {
	m.WalkPushButton.Font = font
	return m
}

func (m *MyPushButton) MaxSize(size Size) *MyPushButton {
	m.WalkPushButton.MaxSize = size
	return m
}

func (m *MyPushButton) MinSize(size Size) *MyPushButton {
	m.WalkPushButton.MinSize = size
	return m
}

func (m *MyPushButton) Name(name string) *MyPushButton {
	m.WalkPushButton.Name = name
	return m
}

func (m *MyPushButton) Persiitent(persistent bool) *MyPushButton {
	m.WalkPushButton.Persistent = persistent
	return m
}

func (m *MyPushButton) RightToLeftReading(rl bool) *MyPushButton {
	m.WalkPushButton.RightToLeftReading = rl
	return m
}

func (m *MyPushButton) ToolTipText(tip string) *MyPushButton {
	m.WalkPushButton.ToolTipText = tip
	return m
}

func (m *MyPushButton) Visible(visible bool) *MyPushButton {
	m.WalkPushButton.Visible = visible
	return m
}

func (m *MyPushButton) AlwaysConsumeSpace(a bool) *MyPushButton {
	m.WalkPushButton.AlwaysConsumeSpace = a
	return m
}

func (m *MyPushButton) Column(column int) *MyPushButton {
	m.WalkPushButton.Column = column
	return m
}

func (m *MyPushButton) Row(row int) *MyPushButton {
	m.WalkPushButton.Row = row
	return m
}

func (m *MyPushButton) RowSpan(rowspan int) *MyPushButton {
	m.WalkPushButton.RowSpan = rowspan
	return m
}

func (m *MyPushButton) StretchFactor(sf int) *MyPushButton {
	m.WalkPushButton.StretchFactor = sf
	return m
}

func (m *MyPushButton) Image(image string) *MyPushButton {
	m.WalkPushButton.Image = image
	return m
}

func (m *MyPushButton) Text(text string) *MyPushButton {
	m.WalkPushButton.Text = text
	return m
}

func (m *MyPushButton) ImageAboveText(above bool) *MyPushButton {
	m.WalkPushButton.ImageAboveText = above
	return m
}

func (m *MyPushButton) Create() PushButton {
	return m.WalkPushButton
}
