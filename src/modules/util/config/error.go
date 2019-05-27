package config

import (
	"fmt"
	"time"
)

type ConfigError struct {
	Time    time.Time
	Message string
	Info    string
}

func NewError(t time.Time, msg string, info string) *ConfigError {
	return &ConfigError{
		Time:    t,
		Message: msg,
		Info:    info,
	}
}

func (c *ConfigError) Error() string {
	return fmt.Sprintf("[%v] - message: %v | info: %v \r\n", c.Time, c.Message, c.Info)
}
