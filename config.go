package main

import (
	"modules/config"
)

var (
	configfiles = map[string]string{"picturesize": "picture_size.json", "picturecate": "picture_cate.json"}
	configs     map[string]*config.Config
)

func RegisterConfigs() {
	// configs = config.NewRegister().RegisterAll(configfiles).Run()
}
