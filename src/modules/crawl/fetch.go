package crawl

import (
	"fmt"
	"io/ioutil"
	"log"
	"modules/app"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/WardonNE/util/convert"
	"github.com/axgle/mahonia"
)

type Catalog struct {
	Name, Href, Title string
}

func between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

func LoadCatalogs() []Catalog {
	var catalogs = make([]Catalog, 0)
	var url = app.CrawlConf.RootUrl
	fmt.Println("request url: ", url)
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
	} else {
		fmt.Println("Status Code:", response.StatusCode)
	}
	return catalogs
}

type SizeClass struct {
	Name, Href string
}

func LoadSizeClasses() []SizeClass {
	var sizeclasses = make([]SizeClass, 0)
	var url = app.CrawlConf.RootUrl
	fmt.Println("request url: ", url)
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
	} else {
		fmt.Println("Status Code:", response.StatusCode)
	}
	return sizeclasses
}

type imageListItem struct {
	Name, Url string
}

func LoadImageList(page int) ([]imageListItem, int, int) {
	var imagelist = make([]imageListItem, 0)
	var url = app.CrawlConf.RootUrl
	var (
		activepage = page
		totalpage  int
	)
	if page > 1 {
		url = url + "/index_" + strconv.Itoa(page) + ".htm"
	} else {
		url = url + "/index.htm"
	}
	fmt.Println("request url:", url)
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
		pageconfig := app.CrawlConf.Page
		activepagedom := dom.Find(pageconfig.TargetElement).Eq(pageconfig.TargetElementIndex).Find(pageconfig.ActivePage.TargetElement).Eq(pageconfig.ActivePage.TargetElementIndex)
		activepage, _ = strconv.Atoi(activepagedom.Text())
		totalpagedom := dom.Find(pageconfig.TargetElement).Eq(pageconfig.TargetElementIndex).Find(pageconfig.TotalPage.TargetElement).Eq(pageconfig.TotalPage.TargetElementIndex).Prev()
		totalpagestring := totalpagedom.Text()
		totalpage, _ = strconv.Atoi(totalpagestring)
	} else {
		fmt.Println("Status Code:", response.StatusCode)
	}
	return imagelist, activepage, totalpage
}

func LoadImageByCatalog(url string, page int) ([]imageListItem, int, int) {
	var imagelist = make([]imageListItem, 0)
	if page > 1 {
		url = url + "index_" + strconv.Itoa(page) + ".htm"
	} else {
		url = url + "index.htm"
	}
	fmt.Println("request url:", url)
	var (
		activepage = page
		totalpage  int
	)
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
		pageconfig := app.CrawlConf.Page
		totalpagedom := dom.Find(pageconfig.TargetElement).Eq(pageconfig.TargetElementIndex).Find(pageconfig.TotalPage.TargetElement).Eq(pageconfig.TotalPage.TargetElementIndex).Prev()
		totalpagestring := totalpagedom.Text()
		totalpage, _ = strconv.Atoi(totalpagestring)
	} else {
		fmt.Println("Status Code:", response.StatusCode)
	}
	return imagelist, activepage, totalpage
}

func LoadImageBySize(url string, page int) ([]imageListItem, int, int) {
	var imagelist = make([]imageListItem, 0)
	if page > 1 {
		url = url + "index_" + strconv.Itoa(page) + ".htm"
	} else {
		url = url + "index.htm"
	}
	fmt.Println("request url:", url)
	var (
		activepage = page
		totalpage  int
	)
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
		pageconfig := app.CrawlConf.Page
		totalpagedom := dom.Find(pageconfig.TargetElement).Eq(pageconfig.TargetElementIndex).Find(pageconfig.TotalPage.TargetElement).Eq(pageconfig.TotalPage.TargetElementIndex).Prev()
		totalpagestring := totalpagedom.Text()
		totalpage, _ = strconv.Atoi(totalpagestring)
	} else {
		fmt.Println("Status Code:", response.StatusCode)
	}
	return imagelist, activepage, totalpage
}

func LoadImageBySearchKeyword(keyword string, page int) ([]imageListItem, int, int) {
	var imagelist = make([]imageListItem, 0)
	var (
		activepage = page + 1
		totalpage  int
	)
	params := url.Values{}
	gbkkeyword := mahonia.NewEncoder("GBK").ConvertString(keyword)
	params.Add("keyboard", gbkkeyword)
	params.Add("page", strconv.Itoa(page))
	url := app.CrawlConf.RootUrl + app.CrawlConf.Search.SearchApi + "?" + params.Encode()
	fmt.Println("request url:", url)
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
			log.Panicln("transfer body code error:", err)
		}
		reader := strings.NewReader(bodystring)
		dom, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Panicln("create dom error:", err)
		}
		imagelistconfig := app.CrawlConf.ImageList
		dom.Find(imagelistconfig.TargetElement).Eq(imagelistconfig.TargetElementIndex).Find("li").Each(func(i int, s *goquery.Selection) {
			divElement := s.Find("div.pic_box")
			aElement := s.Find("a")
			if divElement.Length() == 0 {
				href, _ := aElement.Attr("href")
				bElement := aElement.Find("b")
				imagelist = append(imagelist, imageListItem{
					Url:  app.CrawlConf.RootUrl + href,
					Name: bElement.Text(),
				})
			}
		})
		pageconfig := app.CrawlConf.Page
		totalpagedom := dom.Find(pageconfig.TargetElement).Eq(pageconfig.TargetElementIndex).Find(pageconfig.TotalPage.TargetElement).Eq(pageconfig.TotalPage.TargetElementIndex).Prev()
		totalpagestring := totalpagedom.Text()
		totalpage, _ = strconv.Atoi(totalpagestring)
	} else {
		fmt.Println("Status Code:", response.StatusCode)
	}
	return imagelist, activepage, totalpage
}
