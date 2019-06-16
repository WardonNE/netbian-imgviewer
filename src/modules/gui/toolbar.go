package gui

import (
	. "github.com/lxn/walk/declarative"

	"modules/app"
)

func GetToolBarList() Composite {
	return Composite{
		MaxSize:  GetToolBarCompositeSize(),
		Layout:   GetToolBarCompositeHBox(),
		Children: GetToolBarCompositeWidgets(),
	}
}

func GetToolBarCompositeSize() Size {
	return Size{
		app.MainWindowConf.Children.ToolBarComposite.MaxSize.Width,
		app.MainWindowConf.Children.ToolBarComposite.MaxSize.Height,
	}
}

func GetToolBarCompositeHBox() HBox {
	return HBox{}
}

func GetToolBarCompositeWidgets() []Widget {
	return []Widget{
		GetFollowPushButton(),
		GetCancelFollowPushButton(),
		GetDownloadPushButton(),
		GetCatalogLabel(),
		GetCatalogComboBox(),
		GetSizeLabel(),
		GetSizeComboBox(),
		GetSearchLabel(),
		GetSearchLineEdit(),
		GetSearchPushButton(),
		GetPrevPagePushButton(),
		GetNextPagePushButton(),
	}
}

func GetFollowPushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.FollowPushButton
	return PushButton{
		Text: pb.Text,
	}
}

func GetCancelFollowPushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.CancelFollowPushButton
	return PushButton{
		Text: pb.Text,
	}
}

func GetDownloadPushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.DownloadPushButton
	return PushButton{
		Text: pb.Text,
	}
}

func GetCatalogLabel() Label {
	label := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.CatalogLabel
	return Label{
		Text: label.Text,
	}
}

func GetCatalogComboBox() ComboBox {
	cb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.CatalogComboBox
	maxsize := Size{
		cb.MaxSize.Width,
		cb.MaxSize.Height,
	}
	minsize := Size{
		cb.MinSize.Width,
		cb.MinSize.Height,
	}
	return ComboBox{
		MaxSize: maxsize,
		MinSize: minsize,
		// Model: GetCatalogsComboBoxModel()
	}
}

func GetSizeLabel() Label {
	label := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.SizeLabel
	return Label{
		Text: label.Text,
	}
}

func GetSizeComboBox() ComboBox {
	cb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.SizeComboBox
	maxsize := Size{
		cb.MaxSize.Width,
		cb.MaxSize.Height,
	}
	minsize := Size{
		cb.MinSize.Width,
		cb.MinSize.Height,
	}
	return ComboBox{
		MaxSize: maxsize,
		MinSize: minsize,
	}
}

func GetSearchLabel() Label {
	label := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.SearchLabel
	return Label{
		Text: label.Text,
	}
}

func GetSearchLineEdit() LineEdit {
	le := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.LineEdit
	maxsize := Size{
		le.MaxSize.Width,
		le.MaxSize.Height,
	}
	minsize := Size{
		le.MinSize.Width,
		le.MinSize.Height,
	}
	return LineEdit{
		MaxSize: maxsize,
		MinSize: minsize,
	}
}

func GetSearchPushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.SearchPushButton
	return PushButton{
		Text: pb.Text,
	}
}

func GetPrevPagePushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.PrevPagePushButton
	return PushButton{
		Text: pb.Text,
	}
}

func GetNextPagePushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.NextPagePushButton
	return PushButton{
		Text: pb.Text,
	}
}

// func GetCatalogsComboBoxModel() {

// }
