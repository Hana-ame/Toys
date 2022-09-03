package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type BulletinMetaData struct {
	id       int
	name     string
	describe string
}

type BubbleMetaData struct {
	Src       string `json:"imgSrc"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Timestamp string `json:"timestamp"`
}
type S2d struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
type Bubble struct {
	Id int `json:"id"`
	// title      string `json:"title,omitempty"`
	// author     string `json:"author"`
	// timestamp  string `json:"timestamp"`
	Meta       BubbleMetaData `json:"metaData"`
	Content    string         `json:"content"`
	Background string         `json:"background"`
	Pos        S2d            `json:"pos"`
	Size       S2d            `json:"size"`
	Scale      float32        `json:"scale"`
}

func main() {
	go DaoInit()
	time.Sleep(3 * time.Second)

	http.HandleFunc("/bb/api/bubble", apiBubble)
	http.ListenAndServe("127.0.99.1:8080", nil)
}

func apiBubble(w http.ResponseWriter, r *http.Request) {
	// 测试用正式版删掉
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// GET
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	apiBubbleFuncs := make([]func(http.ResponseWriter, *Bubble, string) string, 0)
	apiBubbleFuncs = append(apiBubbleFuncs, BubblePost)
	apiBubbleFuncs = append(apiBubbleFuncs, BubbleResp)

	resJson, err := getRequsetJson(r) //;fmt.Println(resJson)

	status := "normal"
	if err != nil {
		status = "task-resp"
	}

	for _, f := range apiBubbleFuncs {
		status = f(w, &resJson, status)
		if status == "return" {
			break
		}
	}

	return
}

func BubblePost(w http.ResponseWriter, resJson *Bubble, status string) string {
	if status != "normal" && status != "task-post" {
		return status
	}

	InsertBubble(*resJson)

	return "normal"
}
func BubbleResp(w http.ResponseWriter, resJson *Bubble, status string) string {
	if status != "normal" && status != "task-resp" {
		return status
	}
	respJson, err := json.Marshal(QueryBubbles(0, "test"))
	if err != nil {
		return "error"
	}
	w.Write((respJson))
	return "normal"
}

func getRequsetJson(r *http.Request) (d Bubble, err error) {

	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// log.Fatal(err)
		return
	}
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
