package fetcher

import (
	"github.com/PuerkitoBio/goquery"

	"bufio"
	"fmt"
	"strings"
	"time"
)

type Fetcher struct {
}

func NewFetcher() *Fetcher {
	return &Fetcher{}
}

func (f *Fetcher) FetchByUrl(url string) *goquery.Document {
	dom, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(NewError(time.Now(), err.Error(), getErrorInfo("fetch", "createnewdom")))
	}
	return dom
}

func (f *Fetcher) FetchByHtmlString(htmlstring string) *goquery.Document {
	reader := bufio.NewReader(strings.NewReader(htmlstring))
	dom, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(NewError(time.Now(), err.Error(), getErrorInfo("fetch", "createnewdom")))
	}
	return dom
}
