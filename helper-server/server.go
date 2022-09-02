package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

var mutex sync.Mutex

type Node map[string]int64

func (n Node) add(addr string) {
	n[addr] = time.Now().UnixNano()
}

type Site map[string]Node

func (s Site) add(path string, addr string) {
	mutex.Lock()
	defer mutex.Unlock()
	// fmt.Println(s[path])
	if s[path] == nil {
		// fmt.Println(s[path])
		s[path] = make(Node, 1)
		// fmt.Println(s[path])
	}
	// fmt.Println(s[path])
	s[path].add(addr)
}
func (s Site) bytes(path string) []byte {
	mutex.Lock()
	defer mutex.Unlock()
	if s[path] == nil {
		return []byte(`{}`)
	}
	b, err := json.Marshal(s[path])
	if err != nil {
		return []byte(`{}`)
	}
	return b
}

var site Site

func main() {
	site = Site{} // 为啥之前没事
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			server_get(w, r)
		} else if r.Method == "POST" {
			server_post(w, r)
		}
	})

	log.Fatal(http.ListenAndServe("127.0.233.0:8080", nil))
}

func server_get(w http.ResponseWriter, r *http.Request) {
	w.Write(site.bytes(r.URL.Path))
}

func server_post(w http.ResponseWriter, r *http.Request) {
	var s string
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	site.add(r.URL.Path, s)
	server_get(w, r)
}
