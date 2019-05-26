package config

import (
	"os"
	"path/filepath"
	"time"
)

func init() {
	createErrorInfo()
	setConfigDirName()
}

func setConfigDirName() {
	exeName, err := os.Executable()
	if err != nil {
		panic(NewError(time.Now(), err.Error(), getErrorInfo("init", "exename")))
	}
	p := filepath.Dir(exeName)
	configDirName = p + "/../conf/"
}
