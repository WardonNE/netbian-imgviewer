package config

import (
	"io/ioutil"
	"time"

	"github.com/sipt/GoJsoner"
)

type Loader struct {
	configPath string
	content    string
}

func NewLoader(file string) *Loader {
	return &Loader{
		configPath: file,
	}
}

func (l *Loader) Load() *Loader {
	b, err := ioutil.ReadFile(l.configPath)
	if err != nil {
		panic(NewError(time.Now(), err.Error(), getErrorInfo("load", "read")))
	}
	l.content = string(b)
	return l
}

func (l *Loader) Discard() string {
	s, err := GoJsoner.Discard(l.content)
	if err != nil {
		panic(NewError(time.Now(), err.Error(), getErrorInfo("load", "discard")))
	}
	return s
}
