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
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/WardonNE/util/convert"

	"netbian-imgviewer/src/modules/app"
)

var tmpdir, downloaddir, favoritedir string

var settings app.SettingsConfig

var downloadconfig app.DownloadConfig

func init() {
	downloadconfig = *app.DownloadConf
	exepath, err := os.Executable()
	if err != nil {
		log.Panicln("init error:", err)
	}
	binpath := filepath.Dir(exepath)
	tmpdir = binpath + "\\" + downloadconfig.TmpDir

	settings = *app.Settings
	downloaddir = binpath + "\\" + settings.DownloadDir
	favoritedir = binpath + "\\" + settings.FavoriteDir
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

func DownloadIntoDataDir(url, title string) {
	filename := downloaddir + "\\" + Md5(title) + ".jpg"
	exist, _ := fileExist(filename)
	if exist {
		fmt.Println("this image has been downloaded")
	}
	tmpfilename := tmpdir + "\\" + Md5(title) + ".jpg"
	exist, _ = fileExist(tmpfilename)
	if !exist {
		DownloadImage(url, title)
	}
	tmpf, err := os.Open(tmpfilename)
	if err != nil {
		log.Panicln("open tmp file error: ", err)
	}
	defer tmpf.Close()
	newf, err := os.Create(filename)
	if err != nil {
		log.Panicln("open new file error: ", err)
	}
	defer newf.Close()
	io.Copy(newf, tmpf)
}

func AddFavorite(url, title string) (int, error) {
	favoriterecordfile := favoritedir + "\\favorite.xml"
	exist, _ := fileExist(favoriterecordfile)
	if !exist {
		f, err := os.Create(favoriterecordfile)
		if err != nil {
			log.Panicln("create favorite record file error: ", err)
		}
		defer f.Close()
		xmlstring := "<Image href=\"" + url + "\" title=\"" + title + "\"></Image>\r\n"
		_, err = f.WriteString(xmlstring)
		if err != nil {
			return -1, err
		}
		return 1, nil
	} else {
		xmlstring := "<Image href=\"" + url + "\" title=\"" + title + "\"></Image>"
		f, err := os.OpenFile(favoriterecordfile, os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			log.Panicln("open favoriterecordfile error: ", err)
		}
		defer f.Close()
		filecontent, err := ioutil.ReadAll(f)
		if err != nil {
			log.Panicln("read file content error: ", err)
		}
		reg := regexp.MustCompile(xmlstring)
		str := reg.FindAllString(string(filecontent), 1)
		fmt.Println("match string: ", str)
		fmt.Println("file content: ", string(filecontent))
		if len(str) > 0 {
			// walk.MsgBox(mw, "Add Favorite", "this image has been added into favorite list", walk.MsgBoxIconInformation)
			return -2, nil
		}
		_, err = f.WriteString(xmlstring + "\r\n")
		if err != nil {
			return -1, err
		}
		return 1, nil
	}
}

func RemoveFavorite(url, title string) {

}
