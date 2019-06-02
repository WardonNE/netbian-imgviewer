package app

import (
	"log"

	"github.com/WardonNE/util/config"
)

type MainWindowConfig struct {
	Profiles struct {
		Title, Version, Author string
	}
	MainWindowSize struct {
		Size struct {
			Width, Height int
		}
		MinSize struct {
			Width, Height int
		}
		MaxSize struct {
			Width, Height int
		}
	}
}

var MainWindowConf = &MainWindowConfig{}

func init() {
	_, err := config.LoadConfigFile("mainwindows.xml", MainWindowConf)
	if err != nil {
		log.Panicln("Init Module App Err: ", err)
	}
}
