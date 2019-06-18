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
		AssignTo: &mw.searchle,
		MaxSize:  maxsize,
		MinSize:  minsize,
	}
}

func GetSearchPushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.SearchPushButton
	return PushButton{
		AssignTo:  &mw.searchpb,
		Text:      pb.Text,
		OnClicked: searchPbOnClick,
	}
}

func GetPrevPagePushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.PrevPagePushButton
	return PushButton{
		AssignTo:  &mw.prevpagepb,
		Text:      pb.Text,
		OnClicked: prevPagePbOnClicked,
	}
}

func GetNextPagePushButton() PushButton {
	pb := app.MainWindowConf.Children.ToolBarComposite.Children.Widget.NextPagePushButton
	return PushButton{
		AssignTo:  &mw.nextpagepb,
		Text:      pb.Text,
		OnClicked: nextPagePbOnClicked,
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
	oldcurrentindex := mw.sizecbcurrentindex
	if i < 0 || i == oldcurrentindex {
		return
	}
	mw.catalogcbcurrentindex = i
	activeItem := &mw.catalogcbmodel.items[i]
	mw.activepage = 1
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

func (mw *MyMainWindow) sizeCbModelOnCurrentIndexChanged() {
	i := mw.sizecb.CurrentIndex()
	fmt.Println("Current Index(Size Combo Box): ", i)
	oldcurrentindex := mw.sizecbcurrentindex
	if i < 0 || i == oldcurrentindex {
		return
	}
	mw.sizecbcurrentindex = i
	activeItem := mw.sizecbmodel.items[i]
	mw.activepage = 1
	reloadImageListModelBySize(activeItem.url, 1)
}

func searchPbOnClick() {
	keyword := mw.searchle.Text()
	fmt.Println("Search Keyword:", keyword)
	if keyword == "" {
		return
	}
	mw.searchkeyword = keyword
	mw.activepage = 1
	reloadIamgeListModeBySearchKeyword(keyword, 0)
}

func prevPagePbOnClicked() {
	var prevpage = mw.activepage - 1
	fmt.Println("Active Page: ", mw.activepage, "Prev Page: ", prevpage)
	if prevpage <= 0 {
		prevpage = 1
	}
	keyword := mw.searchle.Text()
	fmt.Println("Search Keyword: ", keyword)
	if keyword != "" {
		reloadIamgeListModeBySearchKeyword(keyword, prevpage-1)
		return
	}
	catalogcurrentindex := mw.catalogcb.CurrentIndex()
	fmt.Println("Catalog Combo Box Current Index: ", catalogcurrentindex)
	if catalogcurrentindex >= 0 {
		i := mw.catalogcb.CurrentIndex()
		activeItem := mw.catalogcbmodel.items[i]
		reloadImageListModelByCatalog(activeItem.url, prevpage)
		return
	}
	sizecurrentindex := mw.sizecb.CurrentIndex()
	fmt.Println("Size Combo Box Current Index: ", sizecurrentindex)
	if sizecurrentindex >= 0 {
		i := mw.sizecb.CurrentIndex()
		activeItem := mw.sizecbmodel.items[i]
		reloadImageListModelBySize(activeItem.url, prevpage)
		return
	}
	reloadImageList(prevpage)
}

func nextPagePbOnClicked() {
	var nextpage = mw.activepage + 1
	fmt.Println("Active Page: ", mw.activepage, "Next Page: ", nextpage)
	if nextpage > mw.totalpage {
		nextpage = mw.totalpage
	}
	keyword := mw.searchle.Text()
	fmt.Println("Search Keyword: ", keyword)
	if keyword != "" {
		reloadIamgeListModeBySearchKeyword(keyword, nextpage-1)
		return
	}
	catalogcurrentindex := mw.catalogcb.CurrentIndex()
	fmt.Println("Catalog Combo Box Current Index: ", catalogcurrentindex)
	if catalogcurrentindex >= 0 {
		i := mw.catalogcb.CurrentIndex()
		activeItem := mw.catalogcbmodel.items[i]
		reloadImageListModelByCatalog(activeItem.url, nextpage)
		return
	}
	sizecurrentindex := mw.sizecb.CurrentIndex()
	fmt.Println("Size Combo Box Current Index: ", sizecurrentindex)
	if sizecurrentindex >= 0 {
		i := mw.sizecb.CurrentIndex()
		activeItem := mw.sizecbmodel.items[i]
		reloadImageListModelBySize(activeItem.url, nextpage)
		return
	}
	reloadImageList(nextpage)
}
