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


```golang

func main() {
	pool := NewPortalPool(5, 5)
	for i := 0; i < 5; i++ {
		p := NewPortal("udp")
		pool.Add(p)
	}
	fmt.Println(pool)
	// time.Sleep(time.Second * 5)
	fmt.Println(pool)
	// time.Sleep(time.Second * 120)

	p := pool.Pick()
	s, err := GetAddr(p.Conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(p.Conn.LocalAddr())
	fmt.Println(pool)

	for i := 0; i < 4; i++ {
		p := pool.Pick()
		var paddr *string
		laddr := "localhost:9999"
		if i == 0 {
			addr := "localhost:10000"
			paddr = &addr
		}
		if i == 2 {
			addr := "localhost:10001"
			paddr = &addr
		}
		p.Set(paddr, &laddr, nil)
	}

	time.Sleep(time.Second * 60)
}

```

server这边大大概没什么问题


```powershell

.\ncat.exe -vlup 10000
.\ncat.exe -vlup 10001


```

debian的会断开
https://superuser.com/questions/1008348/netcat-keep-listening-for-connection-in-debian


PortalServer没啥问题


