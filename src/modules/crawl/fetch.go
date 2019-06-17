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
	var url = app.CrawlConf.RootUrl
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
		defer response.Body.Close()
		if err != nil {
			log.Panicln("read response body error: ", err)
		}
		bodystring, err := convert.NewConverter(string(body), "GBK", "UTF-8").Transfer()
		if err != nil {
			log.Panicln("transfer body code error: ", err)
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
					Href:  url + href,
				})
			}
		})
	}
	return catalogs
}

type SizeClass struct {
	Name, Href string
}

func LoadSizeClasses() []SizeClass {
	var sizeclasses = make([]SizeClass, 0)
	var url = app.CrawlConf.RootUrl
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panicln("create request error:", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln("send request error:", err)
	}
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			log.Panicln("read response body error:", err)
		}
		bodystring, err := convert.NewConverter(string(body), "GBK", "UTF-8").Transfer()
		if err != nil {
			log.Panicln("transfer body code error:", err)
		}
		reader := strings.NewReader(bodystring)
		dom, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Panicln("create dom error:", dom)
		}
		sizeconfig := app.CrawlConf.Size
		dom.Find(sizeconfig.TargetElement).Eq(sizeconfig.TargetElementIndex).Find("a").Each(func(i int, s *goquery.Selection) {
			_, exist := s.Attr("target")
			if !exist {
				href, _ := s.Attr("href")
				sizeclasses = append(sizeclasses, SizeClass{
					Name: s.Text(),
					Href: url + href,
				})
			}
		})
	}
	return sizeclasses
}

type imageListItem struct {
	Name, Url string
}

func LoadImageList(page int) []imageListItem {
	var imagelist = make([]imageListItem, 0)
	var url = app.CrawlConf.RootUrl
	if page > 1 {
		url = url + "/index_" + string(page) + ".htm"
	} else {
		url = url + "/index.htm"
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panicln("create request error:", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln("send request error:", err)
	}
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Panicln("read body error:", err)
		}
		defer response.Body.Close()
		bodystring, err := convert.NewConverter(string(body), "GBK", "UTF-8").Transfer()
		if err != nil {
			log.Panicln("transfer body code error:", err)
		}
		reader := strings.NewReader(bodystring)
		dom, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Panicln("create dom error:", err)
		}
		imagelistconfig := app.CrawlConf.ImageList
		dom.Find(imagelistconfig.TargetElement).Eq(imagelistconfig.TargetElementIndex).Find("li").Each(func(i int, s *goquery.Selection) {
			aElement := s.Find("a")
			href, _ := aElement.Attr("href")
			title, exist := aElement.Attr("title")
			if exist {
				imagelist = append(imagelist, imageListItem{
					Url:  app.CrawlConf.RootUrl + href,
					Name: title,
				})
			}
		})
	}
	return imagelist
}

func LoadImageByCatalog(url string, page int) []imageListItem {
	var imagelist = make([]imageListItem, 0)
	if page > 1 {
		url = url + "index_" + string(page) + ".htm"
	} else {
		url = url + "index.htm"
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panicln("create request error:", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln("send request error:", err)
	}
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Panicln("read response body error:", err)
		}
		defer response.Body.Close()
		bodystring, err := convert.NewConverter(string(body), "GBK", "UTF-8").Transfer()
		if err != nil {
			log.Panicln("transfer body code error:")
		}
		reader := strings.NewReader(bodystring)
		dom, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Panicln("create dom error: ", err)
		}
		imagelistconfig := app.CrawlConf.ImageList
		dom.Find(imagelistconfig.TargetElement).Eq(imagelistconfig.TargetElementIndex).Find("li").Each(func(i int, s *goquery.Selection) {
			aElement := s.Find("a")
			href, _ := aElement.Attr("href")
			title, exist := aElement.Attr("title")
			if exist {
				imagelist = append(imagelist, imageListItem{
					Url:  app.CrawlConf.RootUrl + href,
					Name: title,
				})
			}
		})
	}
	return imagelist
}

func LoadImageBySize(url string, page int) []imageListItem {
	var imagelist = make([]imageListItem, 0)
	if page > 1 {
		url = url + "index_" + string(page) + ".htm"
	} else {
		url = url + "index.htm"
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panicln("create request error: ", request)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panicln("send request error:", err)
	}
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Panicln("read response body error:", err)
		}
		bodystring, err := convert.NewConverter(string(body), "GBK", "UTF-8").Transfer()
		if err != nil {
			log.Panicln("transfer body code error:", err)
		}
		reader := strings.NewReader(bodystring)
		dom, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Panicln("create dom error:", err)
		}
		imagelistconfig := app.CrawlConf.ImageList
		dom.Find(imagelistconfig.TargetElement).Eq(imagelistconfig.TargetElementIndex).Find("li").Each(func(i int, s *goquery.Selection) {
			aElement := s.Find("a")
			href, _ := aElement.Attr("href")
			title, exist := aElement.Attr("title")
			if exist {
				imagelist = append(imagelist, imageListItem{
					Url:  app.CrawlConf.RootUrl + href,
					Name: title,
				})
			}
		})
	}
	return imagelist
}
