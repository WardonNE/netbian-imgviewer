package crawl

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/WardonNE/util/convert"

	"modules/app"
)

var tmpdir string

var downloadconfig app.DownloadConfig

func init() {
	downloadconfig = *app.DownloadConf
	exepath, err := os.Executable()
	if err != nil {
		log.Panicln("init error:", err)
	}
	binpath := filepath.Dir(exepath)
	tmpdir = binpath + "/" + downloadconfig.TmpDir
}

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func fileExist(path string) (bool, error) {
	_, err := os.Open(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return false, nil
	}
	return false, err
}

func DownloadImage(url, title string) string {
	filename := tmpdir + "\\" + Md5(title) + ".jpg"
	exist, _ := fileExist(filename)
	if !exist {
		url := strings.Replace(url, ".htm", "-1920x1080.htm", 1)
		fmt.Println("request url: ", url)
		response, err := http.Get(url)
		if err != nil {
			log.Panicln("send get request error: ", err)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Panicln("read response body error:", err)
		}
		defer response.Body.Close()
		bodystring, err := convert.NewConverter(string(body), "GBK", "UTF-8").Transfer()
		if err != nil {
			log.Panicln("transfer body code error: ", err)
		}
		reader := strings.NewReader(bodystring)
		dom, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			log.Panicln("create dom error: ", err)
		}
		href, exist := dom.Find(downloadconfig.DownloadUrlElement.TargetElement).Eq(downloadconfig.DownloadUrlElement.TargetElementIndex).Attr("href")
		if exist {
			resp, err := http.Get(href)
			if err != nil {
				log.Panicln("send get request error: ", err)
			}
			f, err := os.Create(filename)
			if err != nil {
				log.Panicln("create file handle error: ", err)
			}
			io.Copy(f, resp.Body)
			defer f.Close()
			return filename
		} else {
			fmt.Println("href not exist")
			return ""
		}
	} else {
		fmt.Println("image exist")
		return filename
	}
	return ""
}
