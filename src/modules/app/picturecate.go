package app

import (
	"modules/util/config"

	"fmt"
	"sync/atomic"
	"time"
)

type PictureCateFetcherConfig struct {
	Element string
	Index   int
	Url     string
	Debug   bool
}

type PictureCateFetcherConfigMgr struct {
	ConfigValues atomic.Value
}

var (
	PictureCateFetcherConf    = &PictureCateFetcherConfig{}
	PictureCateFetcherConfMgr = &PictureCateFetcherConfigMgr{}
)

func init() {
	c := config.NewRegister().Register("picture_cate.json").Load()
	c.AddNotifyer(PictureCateFetcherConfMgr)

	PictureCateFetcherConf.Element = c.GetString("element")
	PictureCateFetcherConf.Index = c.GetInt("index")
	PictureCateFetcherConf.Url = c.GetString("url")
	PictureCateFetcherConf.Debug = c.GetBoolean("debug")

	PictureCateFetcherConfMgr.ConfigValues.Store(PictureCateFetcherConf)

	fmt.Printf("[%v] - Init Picture Size Fetcher Config Done! \r\n", time.Now())
}

func (p *PictureCateFetcherConfigMgr) Callback(c *config.Config) {
	conf := &PictureCateFetcherConfig{}
	conf.Element = c.GetString("element")
	conf.Index = c.GetInt("index")
	conf.Url = c.GetString("url")
	conf.Debug = c.GetBoolean("debug")

	PictureCateFetcherConfMgr.ConfigValues.Store(conf)

	fmt.Printf("[%v] - Reload Picture Size Fetcher Config Done! \r\n", time.Now())
}
