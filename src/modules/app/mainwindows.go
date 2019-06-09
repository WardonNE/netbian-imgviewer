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
	MenuList struct {
		AboutMenu struct {
			Text  string
			Items struct {
				MenuList struct {
					Separator struct{}
					Action    struct {
						Text string
					}
				}
			}
		}
	}
	VBox struct {
		MarginsZero bool
	}
	Children struct {
		ToolBarComposite struct {
			MaxSize struct {
				Width, Height int
			}
			HBox     struct{}
			Children struct {
				Widget struct {
					FollowPushButton struct {
						Text string
					}
					CancelFollowPushButton struct {
						Text string
					}
					CatalogLabel struct {
						Text string
					}
					SizeLabel struct {
						Text string
					}
					SearchLabel struct {
						Text string
					}
					LineEdit struct {
						Text string
					}
					SearchPushButton struct {
						Text string
					}
					PrevPagePushButton struct {
						Text string
					}
					NextPagePushButton struct {
						Text string
					}
				}
			}
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
