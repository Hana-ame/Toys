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



```golang
func main() {
	ps := NewPortalServer("localhost:9999")
	ps.NewPortal()
	ps.NewPortal()
	ps.NewPortal()
	ps.NewPortal()

	var paddr string

	paddr = "localhost:10000"
	ps.ActivePortal(&paddr)
	paddr = "localhost:10001"
	ps.ActivePortal(&paddr)

	time.Sleep(time.Second * 90)
}
```
PortalServer没啥问题




```golang 
// stun-tester.go

func _main() {
	Conn, err := net.ListenPacket("udp", fmt.Sprintf("0.0.0.0:%d", 12321))
	if err != nil {
		log.Fatal("sb")
		return
	}
	// fmt.Println(GetAddr(Conn))

	// addr, err := net.ResolveUDPAddr("udp", "34.145.70.165:12421")
	// if err != nil {
	// 	log.Printf("error : %v", err)
	// 	return
	// }
	// fmt.Println("1")
	// time.Sleep(time.Second * 3)
	// Conn.WriteTo([]byte{0}, addr)
	// Conn.WriteTo([]byte{0}, addr)
	// Conn.WriteTo([]byte{0}, addr)
	// Conn.WriteTo([]byte{0}, addr)
	// Conn.WriteTo([]byte{0}, addr)
	// fmt.Println("2")

	// buffer := make([]byte, 2048)
	// for {
	// 	fmt.Println(3)
	// 	n, addr, err := Conn.ReadFrom(buffer)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	fmt.Printf("packet-received: bytes=%d from=%s\n", n, addr.String())
	// 	// portalproxy.PrintHex(buffer[:n])
	// }
	s, _ := GetAddr(Conn)
	fmt.Println(s)
}