package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

type Parser struct {
	rawData string
}

func NewParser(s string) *Parser {
	return &Parser{rawData: s}
}

func (p *Parser) Parse() map[string]interface{} {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reader := strings.NewReader(p.rawData)
	decoder := json.NewDecoder(reader)
	jsonObject := make(map[string]interface{})
	err := decoder.Decode(&jsonObject)
	if err != nil {
		fmt.Println(NewError(time.Now(), err.Error(), getErrorInfo("parse", "json")))
	}
	return jsonObject
}
