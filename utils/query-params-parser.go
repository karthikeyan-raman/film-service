package utils

import (
	"net/url"
)

//ParseQueryParams --> Elegantly parse queryparams to key value pair
func ParseQueryParams(params url.Values) map[string]string {
	paramsMap := make(map[string]string)
	for key, value := range params {
		paramsMap[key] = value[0]
	}
	if _, hasLimit := paramsMap["limit"]; !hasLimit {
		paramsMap["limit"] = "100"
	}
	if _, hasSkip := paramsMap["skip"]; !hasSkip {
		paramsMap["skip"] = "0"
	}
	return paramsMap
}
