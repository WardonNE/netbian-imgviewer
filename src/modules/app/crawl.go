package app

import (
	"log"

	"github.com/WardonNE/util/config"
)

type CrawlConfig struct {
	RootUrl  string
	Catalogs struct {
		TargetElement      string
		TargetElementIndex int
	}
	Size struct {
		TargetElement      string
		TargetElementIndex int
	}
	ImageList struct {
		TargetElement      string
		TargetElementIndex int
	}
	Search struct {
		SearchApi    string
		SearchResult struct {
			TargetElement      string
			TargetElementIndex int
		}
	}
	Page struct {
		TargetElement      string
		TargetElementIndex int
		ActivePage         struct {
			TargetElement      string
			TargetElementIndex int
		}
		TotalPage struct {
			TargetElement      string
			TargetElementIndex int
		}
	}
}

var CrawlConf = &CrawlConfig{}

func init() {
	_, err := config.LoadConfigFile("crawl.xml", CrawlConf)
	if err != nil {
		log.Panicln("Init Module App Err: ", err)
	}
}
