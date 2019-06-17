package gui

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"modules/app"
	"modules/crawl"
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
		MaxSize: Size{
			label.MaxSize.Width,
			label.MaxSize.Height,
		},
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
		AssignTo:              &mw.catalogcb,
		MaxSize:               maxsize,
		MinSize:               minsize,
		Model:                 GetCatalogsComboBoxModel(),
		OnCurrentIndexChanged: mw.catalogCbModelOnCurrentIndexChanged,
	}
}

func GetSizeLabel() Label {
	label := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.SizeLabel
	return Label{
		Text: label.Text,
		MaxSize: Size{
			label.MaxSize.Width,
			label.MaxSize.Height,
		},
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
		AssignTo:              &mw.sizecb,
		MaxSize:               maxsize,
		MinSize:               minsize,
		Model:                 GetSizeComboBoxModel(),
		OnCurrentIndexChanged: mw.sizeCbModelOnCurrentIndexChanged,
		// CurrentIndex:          0,
	}
}

func GetSearchLabel() Label {
	label := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.SearchLabel
	return Label{
		Text: label.Text,
		MaxSize: Size{
			label.MaxSize.Width,
			label.MaxSize.Height,
		},
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

type catalogComboBoxModel struct {
	walk.ListModelBase
	items []catalogItem
}

type catalogItem struct {
	url, title string
}

func (c *catalogComboBoxModel) ItemCount() int {
	return len(c.items)
}

func (c *catalogComboBoxModel) Value(index int) interface{} {
	return c.items[index].title
}

func GetCatalogsComboBoxModel() *catalogComboBoxModel {
	catalogs := crawl.LoadCatalogs()
	model := &catalogComboBoxModel{
		items: make([]catalogItem, len(catalogs)),
	}
	for key, catalog := range catalogs {
		model.items[key] = catalogItem{
			catalog.Href,
			catalog.Name,
		}
	}
	mw.catalogcbmodel = model
	return model
}

func (mw *MyMainWindow) catalogCbModelOnCurrentIndexChanged() {
	i := mw.catalogcb.CurrentIndex()
	fmt.Println("Current Index(Catalog Combo Box):", i)
	if i < 0 {
		return
	}
	activeItem := &mw.catalogcbmodel.items[i]
	reloadImageListModelByCatalog(activeItem.url, 1)
}

type sizeComboBoxModel struct {
	walk.ListModelBase
	items []sizeItem
}

type sizeItem struct {
	url, title string
}

func GetSizeComboBoxModel() *sizeComboBoxModel {
	sizeclasses := crawl.LoadSizeClasses()
	model := &sizeComboBoxModel{
		items: make([]sizeItem, len(sizeclasses)),
	}
	for key, sizeclass := range sizeclasses {
		model.items[key] = sizeItem{
			url:   sizeclass.Href,
			title: sizeclass.Name,
		}
	}
	mw.sizecbmodel = model
	return model
}

func (s *sizeComboBoxModel) ItemCount() int {
	return len(s.items)
}

func (s *sizeComboBoxModel) Value(index int) interface{} {
	return s.items[index].title
}

func (m *MyMainWindow) sizeCbModelOnCurrentIndexChanged() {
	i := m.sizecb.CurrentIndex()
	fmt.Println("Current Index(Size Combo Box): ", i)
	if i < 0 {
		return
	}
	activeItem := m.sizecbmodel.items[i]
	reloadImageListModelBySize(activeItem.url, 1)
}
