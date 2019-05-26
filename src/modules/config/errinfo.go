package config

import ()

var errinfo map[string]map[string]string

func createErrorInfo() {
	parseinfo := make(map[string]string)
	parseinfo["json"] = "Json parse failed!"

	fileinfo := make(map[string]string)
	fileinfo["open"] = "Open config file failed!"
	fileinfo["stat"] = "Stat config file failed!"

	loadinfo := make(map[string]string)
	loadinfo["read"] = "Read config file failed!"
	loadinfo["discard"] = "Discard content error!"

	initinfo := make(map[string]string)
	initinfo["exename"] = "Get exe name failed!"
	initinfo["exedir"] = "Get exe dir failed!"

	errinfo = make(map[string]map[string]string)
	errinfo["parse"] = parseinfo
	errinfo["file"] = fileinfo
	errinfo["init"] = initinfo
	errinfo["load"] = loadinfo
}

func getErrorInfo(m, k string) string {
	return errinfo[m][k]
}
