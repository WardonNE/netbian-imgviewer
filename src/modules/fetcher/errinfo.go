package fetcher

var (
	errinfo map[string]map[string]string
)

func CreateErrorInfo() {
	errinfo = make(map[string]map[string]string)

	fetchinfo := make(map[string]string)
	fetchinfo["nourl"] = "Url can not be null!"
	fetchinfo["request"] = "Send request failed!"
	fetchinfo["statuscode"] = "HTTP Status Code Is Not 200!"
	fetchinfo["createnewdom"] = "Create new html domcument failed!"
	errinfo["fetch"] = fetchinfo

	// httpinfo := make(map[string]string)
}

func getErrorInfo(m, k string) string {
	return errinfo[m][k]
}
