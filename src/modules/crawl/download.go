package crawl

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

func DownloadImage(url, title string) {
	filename := tmpdir + "/" + Md5(title)
	fmt.Println("filename:", filename)
	// exist, err := fileExist(filename)
	// if err != nil {
	// 	log.Panicln("check file exist error:", err)
	// }
	// if !exist {
	response, err := http.Get(url)
	if err != nil {
		log.Panicln("send get request error:", err)
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Panicln("create file error: ", err)
	}
	io.Copy(f, response.Body)
	// }
}
