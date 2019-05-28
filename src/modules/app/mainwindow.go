package app

import (
	"fmt"
	"sync/atomic"
	"time"

	"modules/util/config"
)

var (
	MainWindowConf    *MainWindowConfig
	MainWindowConfMgr = &MainWindowConfigMgr{}
)

type MainWindowConfig struct {
	Title     string
	Width     int
	Height    int
	MinWidth  int
	MinHeight int
	MaxWidth  int
	MaxHeight int
	Enabled   bool
	debug     bool
}

type MainWindowConfigMgr struct {
	ConfigValues atomic.Value
}

func init() {
	conf := config.NewRegister().Register("mainwindow.json").Load()
	MainWindowConf = &MainWindowConfig{
		Title:     conf.GetString("title"),
		Width:     conf.GetInt("width"),
		Height:    conf.GetInt("height"),
		MinWidth:  conf.GetInt("min_width"),
		MinHeight: conf.GetInt("min_height"),
		MaxWidth:  conf.GetInt("max_width"),
		MaxHeight: conf.GetInt("max_height"),
		Enabled:   conf.GetBoolean("enabled"),
		debug:     conf.GetBoolean("debug"),
	}

	MainWindowConfMgr.ConfigValues.Store(MainWindowConf)
	fmt.Printf("[%v] - Init MainWindow Config Done! \r\n", time.Now())
}

func (m *MainWindowConfigMgr) Callback(c *config.Config) {
	conf := &MainWindowConfig{
		Title:     c.GetString("title"),
		Width:     c.GetInt("width"),
		Height:    c.GetInt("height"),
		MinWidth:  c.GetInt("min_width"),
		MinHeight: c.GetInt("min_height"),
		MaxWidth:  c.GetInt("max_width"),
		MaxHeight: c.GetInt("max_height"),
		Enabled:   c.GetBoolean("enabled"),
		debug:     c.GetBoolean("debug"),
	}

	MainWindowConfMgr.ConfigValues.Store(conf)

	fmt.Printf("[%v] - Reload MainWindow Config Done! \r\n", time.Now())

}
