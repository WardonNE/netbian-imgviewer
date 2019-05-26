package config

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	configDirName string
	ConfigList    = make(map[string]*Config)
)

type Config struct {
	filename       string
	data           map[string]interface{}
	lastModifyTime int64
	notifyList     []Notifyer
	sync.RWMutex
}

func NewConfig(filename string) *Config {

	f, err := os.Open(configDirName + filename)
	if err != nil {
		panic(NewError(time.Now(), err.Error(), getErrorInfo("file", "open")))
	}

	fi, err := f.Stat()
	if err != nil {
		panic(NewError(time.Now(), err.Error(), getErrorInfo("file", "stat")))
	}

	conf := &Config{
		filename:       filename,
		data:           make(map[string]interface{}),
		lastModifyTime: fi.ModTime().Unix(),
	}
	s := NewLoader(configDirName + filename).Load().Discard()

	conf.Lock()
	conf.data = NewParser(s).Parse()
	conf.Unlock()
	if conf.GetBoolean("debug") {
		fmt.Printf("[%v] - Config: %v \r\n", time.Now(), conf.data)
	}
	go conf.Reload()
	return conf
}

func (c *Config) Reload() {
	ticker := time.NewTicker(time.Second * 5)
	for _ = range ticker.C {
		func() {

			configFile := configDirName + c.filename

			f, err := os.Open(configFile)
			if err != nil {
				panic(NewError(time.Now(), err.Error(), getErrorInfo("file", "open")))
			}

			fi, err := f.Stat()
			if err != nil {
				panic(NewError(time.Now(), err.Error(), getErrorInfo("file", "stat")))
			}

			var currentModifyTime = fi.ModTime().Unix()

			if currentModifyTime > c.lastModifyTime {
				s := NewLoader(configFile).Load().Discard()
				c.Lock()
				c.data = NewParser(s).Parse()
				c.Unlock()
				c.lastModifyTime = currentModifyTime
				if c.GetBoolean("debug") {
					fmt.Printf("[%v] - Config: %v \r\n", time.Now(), c.data)
				}
				for _, n := range c.notifyList {
					n.Callback(c)
				}
			}

		}()
	}
}

func (c *Config) AddNotifyer(n Notifyer) {
	c.notifyList = append(c.notifyList, n)
}

func (c *Config) AddNotifyerAll(ns []Notifyer) {
	c.notifyList = append(c.notifyList, ns...)
}

func (c *Config) GetString(k string) string {
	return c.data[k].(string)
}

func (c *Config) GetBoolean(k string) bool {
	return c.data[k].(bool)
}

func (c *Config) GetInt(k string) int {
	return int(c.data[k].(float64))
}

func (c *Config) GetFloat(k string) float64 {
	return c.data[k].(float64)
}
