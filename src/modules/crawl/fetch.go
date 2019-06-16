package crawl

import (
	"io/ioutil"
	"log"
	"modules/app"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/WardonNE/util/convert"
)

type Catalog struct {
	Name, Href, Title string
}

func LoadCatalogs() []Catalog {
	var catalogs = make([]Catalog, 0)
	url := app.CrawlConf.RootUrl
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panicln("create request error: ", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln("send request error: ", err)
	}
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Panicln("read response body error: ", err)
		}
		bodystring, err := convert.NewConverter(string(body), "GBK", "UTF-8").Transfer()
		if err != nil {
			log.Panicln("parse body error: ", err)
		}
		reader := strings.NewReader(bodystring)
		dom, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Panicln("create dom error: ", err)
		}
		catalogconfig := app.CrawlConf.Catalogs
		dom.Find(catalogconfig.TargetElement).Eq(catalogconfig.TargetElementIndex).Find("a").Each(func(i int, s *goquery.Selection) {
			_, exist := s.Attr("target")
			if !exist {
				title, _ := s.Attr("title")
				href, _ := s.Attr("href")
				catalogs = append(catalogs, Catalog{
					Title: title,
					Name:  s.Text(),
					Href:  href,
				})
			}
		})
	}
	return catalogs
}
