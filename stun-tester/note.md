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

func main() {
	ps := NewPortalServer("localhost:9999")
	ps.NewPortal()
	ps.NewPortal()
	ps.NewPortal()
	ps.NewPortal()

	fmt.Println(ps.Pool)

	var paddr string

	paddr = "localhost:10000"
	ps.ActivePortal(&paddr,ps.LocalAddr,nil)
	paddr = "localhost:10001"
	ps.ActivePortal(&paddr,ps.LocalAddr,nil)

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
```

```golang 
func __1_main() {
	p1 := &Portal{}
	p2 := &Portal{}
	p3 := &Portal{}
	fmt.Println([]*Portal{p1, p2, p3})

	pool := NewPortalPool(1, 1)
	fmt.Println(pool)

	pool.Add(p1)
	fmt.Println(pool)
	pool.Add(p2)
	fmt.Println(pool)
	pool.Add(p3)
	fmt.Println(pool)

	fmt.Println(pool.Pick())
	fmt.Println(pool)
	fmt.Println(pool.Pick())
	fmt.Println(pool)
	fmt.Println(pool.Pick())
	fmt.Println(pool)

	pool.Add(p1)
	fmt.Println(pool)
	pool.Add(p2)
	fmt.Println(pool)
	pool.Add(p3)
	fmt.Println(pool)

	fmt.Println(pool.Pick())
	fmt.Println(pool)
	fmt.Println(pool.Pick())
	fmt.Println(pool)
	fmt.Println(pool.Pick())
	fmt.Println(pool)

	fmt.Println(pool.Pick())
	fmt.Println(pool)
	fmt.Println(pool.Pick())
	fmt.Println(pool)
	fmt.Println(pool.Pick())
	fmt.Println(pool)

	pool.Add(p1)
	fmt.Println(pool)
	fmt.Println(pool.Pick())
	fmt.Println(pool)

}

```



```golang
func main() {
	pc := NewPortalClient("localhost:10000")
	pc.NewPortal()
	pc.NewPortal()
	pc.NewPortal()
	pc.NewPortal()

	fmt.Println(pc.Pool)

	var paddr string = ""

	paddr = "localhost:9999"
	pc.ActivePortal(&paddr, nil, pc.Mux)
	pc.ActivePortal(&paddr, nil, pc.Mux)
	pc.ActivePortal(&paddr, nil, pc.Mux)
	pc.ActivePortal(&paddr, nil, pc.Mux)

	// p := pc.Mux.Pool.m[0]
	// go func() {
	// 	for {
	// 		fmt.Println(p)
	// 		time.Sleep(time.Second * 5)
	// 	}
	// }()
	for {
		fmt.Println(pc.Pool)
		fmt.Println(pc.Mux)
		fmt.Println(pc.Mux.Pool)
		time.Sleep(time.Second * 2)
		fmt.Println("==============")
	}

	time.Sleep(time.Second * 30)
	fmt.Println("==============")

	fmt.Println(pc.Pool)
	fmt.Println(pc.Mux)
	fmt.Println(pc.Mux.Pool)

	fmt.Println("==============")
	time.Sleep(time.Second * 60)
	fmt.Println("==============")

	fmt.Println(pc.Pool)
	fmt.Println(pc.Mux)
	fmt.Println(pc.Mux.Pool)

	fmt.Println("==============")
	time.Sleep(time.Second * 90)
}


```

PortalClient大概没什么问题。


```golang

	// getPool = pc.Pool
	// putPool = pc.Mux.Pool


```



peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer 
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
peer
ActiveClientPortal &{58.39.90.204:48444 0xc00018e010 <nil> <nil> <nil> 0 10 0} &{[<nil> 0xc00029e280 0xc0003066e0 0xc00029e410 0xc000306b90] 1 0 4 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}}
peer
127.0.0.1:59913
ActiveClientPortal &{58.39.90.204:48441 0xc000208008 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> 0xc0003066e0 0xc00029e410 0xc000306b90] 2 0 3 5 10 {0 0}} &{[0xc0003061e0 <nil> <nil> <nil> <nil>] 0 1 1 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:48445 0xc00018e000 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> 0xc000306b90] 4 0 1 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 1 1 0 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:48442 0xc000086008 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 1 1 0 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:48443 0xc000288000 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> 0xc000306b90] 4 0 1 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 1 1 0 5 10 {0 0}}
58.39.90.204:45504
58.39.90.204:45508
58.39.90.204:45505
58.39.90.204:45506
58.39.90.204:45507
58.39.90.204:45510
2022/09/03 07:11:52 peer :  58.39.90.204:48432
2022/09/03 07:11:52 peer :  58.39.90.204:48437
2022/09/03 07:11:52 peer :  58.39.90.204:48433
2022/09/03 07:11:52 peer :  58.39.90.204:48436
2022/09/03 07:11:52 peer :  58.39.90.204:48434
&{[<nil> 0xc00029e280 0xc0003066e0 0xc00029e410 0xc000306b90] 1 0 4 5 10 {0 0}}
&{[0xc00005cb40 0xc0001a2cd0 0xc00029e050 0xc0001a2d70 0xc0003a4000 0xc0001a2e10] 
0 0 6 6 10 {0 0}}
map[127.0.0.1:59913:0xc0003061e0]
===============
&{[<nil> 0xc00029e280 0xc0003066e0 0xc00029e410 0xc000306b90] 1 0 4 5 10 {0 0}}
&{[0xc00005cb40 0xc0001a2cd0 0xc00029e050 0xc0001a2d70 0xc0003a4000 0xc0001a2e10] 
0 0 6 6 10 {0 0}}
===============
&{[<nil> 0xc00029e280 0xc0003066e0 0xc00029e410 0xc000306b90] 1 0 4 5 10 {0 0}}   
&{[0xc00005cb40 0xc0001a2cd0 0xc00029e050 0xc0001a2d70 0xc0003a4000 0xc0001a2e10] 
0 0 6 6 10 {0 0}}
map[127.0.0.1:59913:0xc0003061e0]
===============
&{[<nil> 0xc00029e280 0xc0003066e0 0xc00029e410 0xc000306b90] 1 0 4 5 10 {0 0}}   
&{[0xc00005cb40 0xc0001a2cd0 0xc00029e050 0xc0001a2d70 0xc0003a4000 0xc0001a2e10] 
0 0 6 6 10 {0 0}}
map[127.0.0.1:59913:0xc0003061e0]
===============
PS D:\Workplace\Toys\stun-tester> ^C
PS D:\Workplace\Toys\stun-tester> go build -o client.exe . ; ./client.exe
well, did you worked?
58.39.90.204:45531
58.39.90.204:45528
58.39.90.204:45532
58.39.90.204:45529
58.39.90.204:45533
2022/09/03 07:12:35 [Mux]receive from  127.0.0.1:59913 len= 1
peer 
peer 
ActiveClientPortal &{58.39.90.204:45531 0xc000308000 <nil> <nil> <nil> 0 10 0} &{[<nil> 0xc0003860f0 0xc00022c1e0 0xc000386190 0xc00022c500] 1 0 4 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45528 0xc00020e000 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> 0xc00022c1e0 0xc000386190 0xc00022c500] 2 0 3 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45532 0xc00018e000 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> 0xc000386190 0xc00022c500] 3 0 2 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45529 0xc00018e010 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> 0xc00022c500] 4 0 1 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45533 0xc000086008 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}}
58.39.90.204:45536
58.39.90.204:45534
58.39.90.204:45537
58.39.90.204:45535
58.39.90.204:45539
127.0.0.1:59913
ActiveClientPortal &{58.39.90.204:45536 0xc00020e320 <nil> <nil> <nil> 0 10 0} &{[<nil> 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280] 1 0 4 5 10 {0 0}} &{[<nil> 0xc0003860f0 0xc00022c1e0 0xc000386190 0xc00022c500] 1 0 4 5 10 {0 0}}        
ActiveClientPortal &{58.39.90.204:45534 0xc000086010 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> 0xc00054e140 0xc00054e1e0 0xc00054e280] 2 0 3 5 10 {0 0}} &{[<nil> 0xc0003860f0 0xc00022c1e0 0xc000386190 0xc00022c500] 1 0 4 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45537 0xc000308300 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> 0xc00054e1e0 0xc00054e280] 3 0 2 5 10 {0 0}} &{[<nil> 0xc0003860f0 0xc00022c1e0 0xc000386190 0xc00022c500] 1 0 4 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45535 0xc00018e3e8 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> 0xc00054e280] 4 0 1 5 10 {0 0}} &{[<nil> 0xc0003860f0 0xc00022c1e0 0xc000386190 0xc00022c500] 1 0 4 5 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45539 0xc000562000 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}} &{[<nil> 0xc0003860f0 0xc00022c1e0 0xc000386190 0xc00022c500] 1 0 4 5 10 {0 0}}
2022/09/03 07:12:37 peer :  58.39.90.204:45509
2022/09/03 07:12:37 peer :  58.39.90.204:48447
2022/09/03 07:12:37 peer :  58.39.90.204:45511
2022/09/03 07:12:37 peer :  58.39.90.204:45512
2022/09/03 07:12:37 peer :  58.39.90.204:45513
58.39.90.204:45547
58.39.90.204:45544
58.39.90.204:45550
58.39.90.204:45546
58.39.90.204:45545
2022/09/03 07:12:38 peer :  58.39.90.204:45538
2022/09/03 07:12:38 peer :  58.39.90.204:45543
2022/09/03 07:12:38 peer :  58.39.90.204:45541
2022/09/03 07:12:38 peer :  58.39.90.204:45540
2022/09/03 07:12:38 peer :  58.39.90.204:45542
2022/09/03 07:12:40 [Mux]receive from  127.0.0.1:59913 len= 2
2022/09/03 07:12:40 [Mux]receive from  127.0.0.1:59913 len= 2
&{[0xc000286be0 0xc0003860f0 0xc00022c1e0 0xc000386190 0xc00022c500 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280] 1 1 9 9 10 {0 0}}
&{[0xc000286eb0 0xc000386dc0 0xc000386e60 0xc00005d5e0 0xc000386f50] 0 0 5 5 10 {0 0}}
map[127.0.0.1:59913:0xc00022c140]
===============
2022/09/03 07:12:46 [Mux]receive from  127.0.0.1:50604 len= 2
127.0.0.1:50604
ActiveClientPortal &{58.39.90.204:45547 0xc000006518 <nil> <nil> <nil> 0 10 0} &{[<nil> 0xc000386dc0 0xc000386e60 0xc00005d5e0 0xc000386f50] 1 0 4 5 10 {0 0}} &{[0xc000286be0 <nil> 0xc00022c1e0 0xc000386190 0xc00022c500 0xc000386be0 0xc00054e140 
0xc00054e1e0 0xc00054e280] 2 1 8 9 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45544 0xc000562010 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> 0xc000386e60 0xc00005d5e0 0xc000386f50] 2 0 3 5 10 {0 0}} &{[0xc000286be0 <nil> 0xc00022c1e0 0xc000386190 0xc00022c500 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280] 2 1 8 9 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45550 0xc0004d4000 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> 0xc00005d5e0 0xc000386f50] 3 0 2 5 10 {0 0}} &{[0xc000286be0 <nil> 0xc00022c1e0 0xc000386190 0xc00022c500 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280] 2 1 8 9 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45546 0xc00018e3f8 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> 0xc000386f50] 4 0 1 5 10 {0 0}} &{[0xc000286be0 <nil> 0xc00022c1e0 0xc000386190 0xc00022c500 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280] 2 1 8 9 10 {0 0}}
ActiveClientPortal &{58.39.90.204:45545 0xc000308320 <nil> <nil> <nil> 0 10 0} &{[<nil> <nil> <nil> <nil> <nil>] 0 0 0 5 10 {0 0}} &{[0xc000286be0 <nil> 0xc00022c1e0 0xc000386190 0xc00022c500 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280] 2 1 8 9 10 {0 0}}
58.39.90.204:45557
58.39.90.204:45561
58.39.90.204:45563
58.39.90.204:45558
58.39.90.204:45559
2022/09/03 07:12:47 peer :  58.39.90.204:45551
2022/09/03 07:12:47 peer :  58.39.90.204:45552
2022/09/03 07:12:47 peer :  58.39.90.204:45553
2022/09/03 07:12:47 peer :  58.39.90.204:45554
2022/09/03 07:12:47 peer :  58.39.90.204:45548
2022/09/03 07:12:48 [Mux]receive from  127.0.0.1:50604 len= 2
2022/09/03 07:12:51 read udp [::]:53064: use of closed network connection
&{[0xc000286be0 0xc000286eb0 0xc000386e60 0xc00005d5e0 0xc000386f50 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280 0xc000386dc0] 5 5 10 10 10 {0 0}}
&{[0xc0004a05f0 0xc00005d900 0xc0004a0690 0xc0004a0730 0xc00005d9a0] 0 0 5 5 10 {0 0}}
map[127.0.0.1:50604:0xc0003860f0]
===============
2022/09/03 07:12:59 read udp [::]:53062: use of closed network connection
&{[0xc000286be0 0xc000286eb0 0xc000386e60 0xc00005d5e0 0xc000386f50 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280 0xc000386dc0] 5 5 10 10 10 {0 0}}
&{[0xc0004a05f0 0xc00005d900 0xc0004a0690 0xc0004a0730 0xc00005d9a0] 0 0 5 5 10 {0 0}}
map[]
===============
&{[0xc000286be0 0xc000286eb0 0xc000386e60 0xc00005d5e0 0xc000386f50 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280 0xc000386dc0] 5 5 10 10 10 {0 0}}
&{[0xc0004a05f0 0xc00005d900 0xc0004a0690 0xc0004a0730 0xc00005d9a0] 0 0 5 5 10 {0 0}}
map[]
===============
&{[0xc000286be0 0xc000286eb0 0xc000386e60 0xc00005d5e0 0xc000386f50 0xc000386be0 0xc00054e140 0xc00054e1e0 0xc00054e280 0xc000386dc0] 5 5 10 10 10 {0 0}}
&{[0xc0004a05f0 0xc00005d900 0xc0004a0690 0xc0004a0730 0xc00005d9a0] 0 0 5 5 10 {0 0}}
map[]
===============