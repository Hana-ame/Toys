单纯测试一下是否是严格的nat类型。
返回一个值时说明不是严格类型，有nat希望。
python版本的不知道放到哪里去了。


```golang

func main() {
	conn, err := net.ListenUDP(`udp`, nil)
	if err != nil {
		log.Fatal(`NewWrapper : `, err.Error())
	}
	go func() {
		time.Sleep(time.Second)
		conn.Close()
	}()
	buf := make([]byte, 100)
	_, _, err = conn.ReadFrom(buf)
	if err != nil {
		log.Println(err)
	}

}

```

会返回已经关闭的conn错误




```bash

nc -uvlp 10000


```