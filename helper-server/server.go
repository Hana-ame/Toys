package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Node struct {
	Server  map[string]int64 `json:"server"`
	Client  map[string]int64 `json:"client"`
	NowTime int64            `json:"nowTime"`
	mu      sync.Mutex
}

func (n *Node) addServer(s string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Server[s] = time.Now().Unix()
}
func (n *Node) addClient(s string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.Client[s] = time.Now().Unix()
}
func (n *Node) bytes() []byte {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.NowTime = time.Now().Unix()
	b, err := json.Marshal(*n)
	if err != nil {
		return []byte(`{}`)
	}
	return b
}

var mu sync.RWMutex
var site map[string]*Node

func main() {
	site = make(map[string]*Node)

	r := mux.NewRouter()
	r.HandleFunc("/{path}/server", server)
	r.HandleFunc("/{path}/client", client)
	r.HandleFunc("/{path}/clear", clear)
	http.ListenAndServe(":8080", r)

}

func clear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path, ok := vars["path"]
	if !ok {
		log.Println("path is missing in parameters")
		log.Println(r.URL.Path)
		return
	}

	mu.Lock()
	delete(site, path)
	mu.Unlock()
}

func server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path, ok := vars["path"]
	if !ok {
		fmt.Println("path is missing in parameters")
		log.Println(r.URL.Path)
		return
	}

	var s string
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.RLock()
	n := site[path]
	mu.RUnlock()

	if n == nil {
		n = &Node{
			Server: make(map[string]int64),
			Client: make(map[string]int64),
		}

		mu.Lock()
		site[path] = n
		mu.Unlock()
	}

	n.addServer(s)

	w.Write(n.bytes())

}
func client(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path, ok := vars["path"]
	if !ok {
		fmt.Println("path is missing in parameters")
		log.Println(r.URL.Path)
		return
	}

	var s string
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.RLock()
	n := site[path]
	mu.RUnlock()

	if n == nil {
		n = &Node{
			Server: make(map[string]int64),
			Client: make(map[string]int64),
		}

		mu.Lock()
		site[path] = n
		mu.Unlock()
	}

	n.addClient(s)

	w.Write(n.bytes())

}
