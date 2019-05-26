package main

import (
	"github.com/PuerkitoBio/goquery"

	"fmt"
	"modules/app"
	"modules/fetcher"
	"modules/util"
)

func main() {
	fetchSize()
	fetchCate()

	for {

	}
}

func init() {
	RegisterConfigs()

}

func fetchSize() {
	conf := app.PictureSizeFetcherConfMgr.ConfigValues.Load().(*app.PictureSizeFetcherConfig)
	dom := fetcher.NewFetcher().FetchByUrl(conf.Url)
	html, _ := goquery.OuterHtml(dom.Find(conf.Element).Eq(conf.Index))
	fmt.Println("HTML: ", util.NewConverter(html, "gbk", "utf8").Translate())
}

func fetchCate() {
	// conf := configs["picturecate"]
	// dom := fetcher.NewFetcher().FetchByUrl(conf.GetString("url"))
	// html, _ := goquery.OuterHtml(dom.Find(conf.GetString("element")))
	// fmt.Println("HTML: ", util.NewConverter(html, "gbk", "utf8").Translate())
}
