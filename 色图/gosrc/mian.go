package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var m []string

func main() {
	http.HandleFunc("/ecchi/api/upload", upload)
	http.HandleFunc("/ecchi/api/today", today)
	http.HandleFunc("/ecchi/api/today/302", today302)
	http.HandleFunc("/ecchi/api/next", next)
	http.HandleFunc("/ecchi/api/clear", clear)
	go func() {
		next(nil, nil)
		time.Sleep(time.Hour * 24)
	}()
	http.ListenAndServe("127.0.233.1:8080", nil)
}
func clear(w http.ResponseWriter, r *http.Request) {
	m = []string{}
}
func next(w http.ResponseWriter, r *http.Request) {
	if len(m) == 0 {
		return
	}
	m = m[1:]
}
func upload(w http.ResponseWriter, r *http.Request) {
	list, err := getRequsetJson(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	m = append(m, list...)
	w.Write([]byte("success"))
}

func today(w http.ResponseWriter, r *http.Request) {
	if len(m) == 0 {
		w.Write([]byte(`没有图片`))
		return
	}
	respJson, err := JSONMarshal(m[:1])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(respJson)
}
func today302(w http.ResponseWriter, r *http.Request) {
	if len(m) == 0 {
		http.Redirect(w, r, `https://pximg.moonchan.xyz/img-original/img/2015/12/24/00/18/35/54191694_p0.jpg`, 302)
		return
	}
	http.Redirect(w, r, m[0], 302)
}

func getRequsetJson(r *http.Request) (d []string, err error) {
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// log.Fatal(err)
		return
	}
	d = make([]string, 256)
	err = json.Unmarshal(resBody, &d)
	if err != nil {
		log.Printf("error decoding sakura response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("sakura response: %q", resBody)
		return
	}
	return
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
