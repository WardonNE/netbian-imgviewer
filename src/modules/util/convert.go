package util

import (
	"github.com/axgle/mahonia"

	"fmt"
)

type Converter struct {
	str        string
	srccode    string
	targetcode string
}

func NewConverter(str, srccode, targetcode string) *Converter {
	return &Converter{
		str:        str,
		srccode:    srccode,
		targetcode: targetcode,
	}
}

func (c *Converter) Translate() string {
	srcCoder := mahonia.NewDecoder(c.srccode)
	srcResult := srcCoder.ConvertString(c.str)
	targetCoder := mahonia.NewDecoder(c.targetcode)
	_, r, err := targetCoder.Translate([]byte(srcResult), true)
	if err != nil {
		fmt.Println(err)
	}
	return string(r)
}
