package main

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

func getResp(r *http.Request) *http.Response {
	// if !shouldProxy(r.URL.Host)
	// fmt.Println(shouldProxy(r.URL.String()))
	trueHost := r.URL.Host
	// fmt.Println(trueHost)
	//////////////////////////////////////
	// fmt.Println("here is Request")
	// 制作请求头
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
	}
	// newUrl:
	// 修改newUrl达到去除SNI
	newUrl := r.URL
	newUrl.Host = getIPbyHost(r.URL.Host) // IP
	newUrl.Scheme = "https"
	// fmt.Println("修改后的newUrl", newUrl.String())

	// DEBUG`
	// 显示修改后URL
	// fmt.Println(newUrl.String())

	// req: newUrl
	// 修改后的请求（改为https，HOST重定向，SNI略）
	req, err := http.NewRequest("GET", newUrl.String(), nil)
	if err != nil {
		fmt.Println(`Error On NewRequest`, "\n")
		return goproxy.NewResponse(r, "", 500, "")
	}
	req.Header = r.Header
	req.Method = r.Method
	req.Body = r.Body   //
	req.Host = trueHost //Host

	// DEBUG:
	// 请求头的内容
	// fmt.Println("请求头：", req.Header)
	referer := req.Header.Get("Referer")
	if referer != "" {
		referer = strings.Replace(referer, "http://", "https://", -1)
		referer = strings.Replace(referer, `http:\/\/`, "https://", -1)
		referer = strings.Replace(referer, `http%3A%2F%2F`, "https%3A%2F%2F", -1)
		req.Header.Set("Referer", referer)
	}

	origin := req.Header.Get("Origin")
	if referer != "" {
		origin = strings.Replace(origin, "http://", "https://", -1)
		origin = strings.Replace(origin, `http:\/\/`, "https://", -1)
		origin = strings.Replace(origin, `http%3A%2F%2F`, "https%3A%2F%2F", -1)
		req.Header.Set("Origin", origin)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(`Error On Do Request`)
		return goproxy.NewResponse(r, "", 500, "")
	}
	defer resp.Body.Close()

	// DEBUG:
	// 响应头的内容
	// fmt.Println("响应头：", resp.Header, "\n")
	// if resp.Header.Get("Set-Cookie") != "" {
	// 	fmt.Println("响应头：(Set-Cookie)", resp.Header.Get("Set-Cookie"), "\n")
	// }
	// if resp.Header.Get("Location") != "" {
	// 	fmt.Println("响应头：(Location)", resp.Header.Get("Location"), "\n")
	// }

	// s:
	// Body内容
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("gzip!", "\n")
			log.Fatal("gzip")
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}
	buf := new(strings.Builder)
	io.Copy(buf, reader)
	s := buf.String()

	// fmt.Println("Response:", s, "\n")
	// writeFile(url.QueryEscape((newUrl.Path)), s)

	// Debug
	// fmt.Println(resp.Header.Get("Content-Type"), resp.StatusCode)
	// 如果是html，则改变一些连接
	contentType := resp.Header.Get("Content-Type")
	if len(contentType) >= 9 {
		if contentType[0:9] == "text/html" {
			s = strings.Replace(s, "https://"+trueHost, "http://"+trueHost, -1)
			s = overWrite(s, trueHost)
			// s = strings.Replace(s, "https://cdnjs.cloudflare.com", "https://cdnjs.moonchan.xyz", -1)
			// s = strings.Replace(s, "https://i.pximg.net", "http://i.pximg.net", -1)
			// s = strings.Replace(s, `https://account.pixiv.net`, "http://account.pixiv.net", -1)
			// s = strings.Replace(s, "https://d.pixiv.org", "http://d.pixiv.org", -1)
			// s = strings.Replace(s, `https:\/\/i.pximg.net`, "http://i.pximg.net", -1)
			// s = strings.Replace(s, `https:\/\/www.pixiv.net`, "http://www.pixiv.net", -1)
			// fmt.Println("Response:", s, "\n")
		} else if len(contentType) >= 16 {
			if contentType[0:16] == "application/json" {
				s = strings.Replace(s, "https://"+trueHost, "http://"+trueHost, -1)
				s = overWrite(s, trueHost)
				// s = strings.Replace(s, "https://i.pximg.net", "http://i.pximg.net", -1)
				// s = strings.Replace(s, `https://account.pixiv.net`, "http://account.pixiv.net", -1)
				// s = strings.Replace(s, "https://d.pixiv.org", "http://d.pixiv.org", -1)
				// s = strings.Replace(s, `https:\/\/i.pximg.net`, "http://i.pximg.net", -1)
				// s = strings.Replace(s, `https:\/\/www.pixiv.net`, "http://www.pixiv.net", -1)
			}
		}
		// writeFile(url.QueryEscape((newUrl.Path)), s)
	}
	// 以上，修改返回网页s的一些内容
	statusCode := resp.StatusCode
	/////////////////////////////

	newResp := goproxy.NewResponse(
		r,
		contentType,
		statusCode,
		s,
	)
	// 尝试搬整个头，还是不行
	// newResp.Header = resp.Header.Clone()
	// newResp.Header.Del("Content-Encoding")

	location := resp.Header.Get("Location")
	if location != "" {
		location = strings.Replace(location, "https://", "http://", -1)
		location = strings.Replace(location, `https:\/\/`, "http://", -1)
		location = strings.Replace(location, `https%3A%2F%2F`, "http%3A%2F%2F", -1)
		newResp.Header.Set("Location", location)
		// fmt.Println("新响应头：(Location)", location, "\n")
	}
	if resp.Header.Get("Set-Cookie") != "" {
		setCookie := resp.Header.Get("Set-Cookie")
		setCookie = strings.Replace(setCookie, `; secure`, "", -1)
		setCookie = strings.Replace(setCookie, `; Secure`, "", -1)
		newResp.Header.Set("Set-Cookie", setCookie)
		// fmt.Println("服务器响应头：", resp.Header, "\n")
		// fmt.Println("对游览器Set-Cookie", newResp.Header.Get("Set-Cookie"), "\n")
	}

	if resp.Header.Get("date") != "" {
		newResp.Header.Set("date", resp.Header.Get("date"))
	}
	if resp.Header.Get("last-modified") != "" {
		newResp.Header.Set("last-modified", resp.Header.Get("last-modified"))
	}
	if resp.Header.Get("cache-control") != "" {
		newResp.Header.Set("cache-control", resp.Header.Get("cache-control"))
	}
	// fmt.Println("新头：", newResp.Header, "\n")
	// newResp.Header.Set("Content-Encoding")
	return newResp
	/////////////////////////////
	// return nil
}
