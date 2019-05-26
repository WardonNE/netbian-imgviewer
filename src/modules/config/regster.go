package config

import (
	"fmt"
	"time"
)

var (
	configFiles = make(map[string]string)
)

type Register struct {
	configFile string
}

func NewRegister() *Register {
	return new(Register)
}

func (r *Register) Register(filename string) *Register {
	fmt.Printf("[%v] - Load New Config File %v \r\n", time.Now(), filename)
	r.configFile = filename
	return r
}

// func (r *Register) RegisterAll(files map[string]string) *Register {
// 	for k, cFile := range files {
// 		r.Register(cFile, k)
// 	}
// 	return r
// }

func (r *Register) Delete(filename string) *Register {
	for k, f := range configFiles {
		if f == filename {
			delete(configFiles, k)
		}
	}
	return r
}

func (r *Register) Load() *Config {
	return NewConfig(r.configFile)
}

// func (r *Register) Run() map[string]*Config {
// 	for k, cFile := range configFiles {
// 		conf := NewConfig(cFile)
// 		ConfigList[k] = conf
// 	}
// 	return ConfigList
// }
