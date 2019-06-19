package app

import (
	"log"

	"github.com/WardonNE/util/config"
)

type DownloadConfig struct {
	TmpDir             string
	Expire             int
	DownloadUrlElement struct {
		TargetElement      string
		TargetElementIndex int
	}
}

var DownloadConf = &DownloadConfig{}

func init() {
	_, err := config.LoadConfigFile("download.xml", DownloadConf)
	if err != nil {
		log.Panicln("Init Module App Err: ", err)
	}
}
