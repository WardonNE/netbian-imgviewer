package app

import (
	"modules/config"

	"fmt"
	"sync/atomic"
	"time"
)

type PictureSizeFetcherConfig struct {
	Element string
	Index   int
	Url     string
	Debug   bool
}

type PictureSizeFetcherConfigMgr struct {
	ConfigValues atomic.Value
}

var (
	PictureSizeFetcherConf    = &PictureSizeFetcherConfig{}
	PictureSizeFetcherConfMgr = &PictureSizeFetcherConfigMgr{}
)

func init() {
	c := config.NewRegister().Register("picture_size.json").Load()
	c.AddNotifyer(PictureSizeFetcherConfMgr)

	PictureSizeFetcherConf.Element = c.GetString("element")
	PictureSizeFetcherConf.Index = c.GetInt("index")
	PictureSizeFetcherConf.Url = c.GetString("url")
	PictureSizeFetcherConf.Debug = c.GetBoolean("debug")

	PictureSizeFetcherConfMgr.ConfigValues.Store(PictureSizeFetcherConf)

	fmt.Printf("[%v] - Init Picture Size Fetcher Config Done! \r\n", time.Now())
}

func (p *PictureSizeFetcherConfigMgr) Callback(c *config.Config) {
	conf := &PictureSizeFetcherConfig{}
	conf.Element = c.GetString("element")
	conf.Index = c.GetInt("index")
	conf.Url = c.GetString("url")
	conf.Debug = c.GetBoolean("debug")

	PictureSizeFetcherConfMgr.ConfigValues.Store(conf)

	fmt.Printf("[%v] - Reload Picture Size Fetcher Config Done! \r\n", time.Now())
}
