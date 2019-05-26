package fetcher

import (
	"fmt"
	"time"
)

type FetcherError struct {
	Time    time.Time
	Message string
	Info    string
}

func NewError(t time.Time, msg string, info string) *FetcherError {
	return &FetcherError{
		Time:    t,
		Message: msg,
		Info:    info,
	}
}

func (f *FetcherError) Error() string {
	return fmt.Sprintf("[%v] - message: %v | info: %v \r\n", f.Time, f.Message, f.Info)
}
