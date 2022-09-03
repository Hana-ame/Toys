package main

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

var host map[string]([]string)

func main() {
	initHost()
	http.HandleFunc("/", nyaa)
	http.ListenAndServe("127.0.0.1:22222", nil)
}

func nyaa(w http.ResponseWriter, r *http.Request) {
	trueHost := "nyaa.si"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
	}

	newUrl := r.URL
	newUrl.Host = getIPbyHost(trueHost) // IP
	newUrl.Scheme = "https"

	req, err := http.NewRequest("GET", newUrl.String(), nil)
	if err != nil {
		fmt.Println(`Error On NewRequest`, "\n")
		// return goproxy.NewResponse(r, "", 500, "")
		return
	}
	req.Header = r.Header
	req.Method = r.Method
	req.Body = r.Body   //
	req.Host = trueHost //Host

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(`Error On Do Request`)
		// return goproxy.NewResponse(r, "", 500, "")
		return
	}
	defer resp.Body.Close()

	// debug
	// fmt.Println("========1632")
	// fmt.Println(resp.Header.Get("Content-Encoding"))

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			// fmt.Println("gzip!", "\n")
			log.Fatal("gzip")
		}
		defer reader.Close()
	case "br":
		r := brotli.NewReader(resp.Body)
		if r == nil {
			log.Fatal("error decoding br response", r)
		}
		// defer reader.Close()
		// statusCode := resp.StatusCode
		contentType := resp.Header.Get("Content-Type")
		w.Header().Set(`Transfer-Encoding`, `br`)
		w.Header().Set("Content-Type", contentType)

		for k, v := range resp.Header {
			// fmt.Println(k)
			if k == `Content-Encoding` || k == `Transfer-Encoding` {
				// fmt.Println(k)
				continue
			}
			w.Header().Set(k, v[0])
		}

		// w.WriteHeader(statusCode)
		io.Copy(w, r)
		// w.WriteHeader(statusCode)
		return
		// 直接结束了
	default:
		reader = resp.Body
	}
	buf := new(strings.Builder)
	io.Copy(buf, reader)
	s := buf.String()

	statusCode := resp.StatusCode
	contentType := resp.Header.Get("Content-Type")
	w.Header().Set("Content-Type", contentType)
	// w.Header().Set(`Transfer-Encoding`, `plain`)

	for k, v := range resp.Header {
		// fmt.Println(k, v)
		if k == `Content-Encoding` || k == `Transfer-Encoding` {
			// fmt.Println(k)
			continue
		}
		w.Header().Set(k, v[0])
		// for _, vv := range v {
		// w.Header().Add(k, vv)
		// }
	}
	// w.Header().Del(`Transfer-Encoding`)
	// w.Header().Set(`Transfer-Encoding`, `plain`)
	if resp.Header.Get("Set-Cookie") != "" {
		setCookie := resp.Header.Get("Set-Cookie")
		setCookie = strings.Replace(setCookie, `; secure`, "", -1)
		setCookie = strings.Replace(setCookie, `; Secure`, "", -1)
		setCookie = strings.Replace(setCookie, `Domain=.nyaa.si;`, "", -1)
		w.Header().Set("Set-Cookie", setCookie)
		// fmt.Println("服务器响应头：", resp.Header, "\n")
		// fmt.Println("对游览器Set-Cookie", newResp.Header.Get("Set-Cookie"), "\n")
	}
	w.WriteHeader(statusCode)
	// fmt.Println(s)
	// fmt.Println(r.URL.Path)
	if r.URL.Path == `/` {
		s = strings.Replace(s, `</nav>`, `</nav><iframe width="100%" height="48" src="https://localhost/nyaa.html">admin@localhost</iframe>`, 1)
	}

	w.Write([]byte(s))

	// w.WriteHeader(statusCode)
}

func getIPbyHost(s string) string {
	return host[s][rand.Intn(len(host[s]))]
}

func initHost() {
	host = make(map[string]([]string), 200)
	host["exhentai.org"] = []string{
		"178.175.129.252",
		"178.175.129.254",
		"178.175.128.252",
		"178.175.128.254",
		"178.175.132.20",
		"178.175.132.22",
	}
	host["nyaa.si"] = []string{
		"185.178.208.182",
	}
	host["sukebei.nyaa.si"] = []string{
		"198.251.89.38",
	}
	host["www.pixiv.net"] = []string{
		"210.140.131.199",
		"210.140.131.219",
		"210.140.131.201",
		"210.140.131.223",
	}
	host["accounts.pixiv.net"] = []string{
		"210.140.131.199",
		"210.140.131.219",
		"210.140.131.201",
		"210.140.131.223",
	}
	host["i.pximg.net"] = []string{
		"210.140.92.141",
		"210.140.92.142",
		"210.140.92.140",
		"210.140.92.144",
		"210.140.92.145",
		"210.140.92.138",
		"210.140.92.146",
		"210.140.92.139",
		"210.140.92.143",
		"210.140.92.147",
	}
	host["d.pixiv.org"] = []string{
		"210.140.131.159",
		"210.140.131.157",
		"210.140.131.158",
	}
	host["www.xvideos.com"] = []string{
		"185.88.181.7",
		"185.88.181.8",
		"185.88.181.9",
		"185.88.181.10",
		"185.88.181.11",
		"185.88.181.2",
		"185.88.181.3",
		"185.88.181.4",
		"185.88.181.5",
		"185.88.181.6",
	}
	host["www.pornhub.com"] = []string{
		"66.254.114.41",
	}
	host["cn.pornhub.com"] = []string{
		"66.254.114.41",
	}
}
