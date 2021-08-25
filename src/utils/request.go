package utils

import (
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

/**
@usage : Retrieve Params and Header from request object
*/
func ProcessRequest(r *http.Request) (http.Header, map[string][]string) {
	headersMap := r.Header
	parmasMap := r.URL.Query()
	return headersMap, parmasMap
}
