package main

import (
	"fmt"
	"time"
)

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
