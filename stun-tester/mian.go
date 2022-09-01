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
	time.Sleep(time.Second * 5)
	fmt.Println(pool)
	time.Sleep(time.Second * 120)

	p := pool.Pick()
	s, err := GetAddr(p.Conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

	fmt.Println(pool)
}
