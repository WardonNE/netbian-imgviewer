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
					DownloadPushButton struct {
						Text string
					}
					CatalogLabel struct {
						Text string
					}
					CatalogComboBox struct {
						MaxSize struct {
							Width, Height int
						}
						MinSize struct {
							Width, Height int
						}
					}
					SizeLabel struct {
						Text string
					}
					SizeComboBox struct {
						MaxSize struct {
							Width, Height int
						}
						MinSize struct {
							Width, Height int
						}
					}
					SearchLabel struct {
						Text string
					}
					LineEdit struct {
						Text    string
						MaxSize struct {
							Width, Height int
						}
						MinSize struct {
							Width, Height int
						}
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
		BodyComposite struct {
			Layout struct {
				Grid struct {
					Columns int
					Spacing int
				}
			}
			ImageListBox struct {
				MaxSize struct {
					Width, Height int
				}
				MinSize struct {
					Width, Height int
				}
			}
			ImageView struct {
				MaxSize struct {
					Width, Height int
				}
				MinSize struct {
					Width, Height int
				}
				Image      string
				Margin     int
				Background struct {
					R, G, B byte
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
