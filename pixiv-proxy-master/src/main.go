//main.go
package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	// init
	// initHost()
	// initPars()
	loadConfig()
	//
	proxy := goproxy.NewProxyHttpServer()
	// 每当HTTP请求时
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			// fmt.Println("HTTP:", r.URL.String())
			if shouldProxy(r.URL.Host) {
				return r, getResp(r)
			}
			return r, nil
		})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", proxy))
}
