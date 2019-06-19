package app

import (
	"log"

	"github.com/WardonNE/util/config"
)

type SettingsConfig struct {
	DownloadDir string
	FavoriteDir string
}

var Settings = &SettingsConfig{}

func init() {
	_, err := config.LoadConfigFile("settings.xml", Settings)
	if err != nil {
		log.Panicln("init module app error: ", err)
	}
}
