package main

import "time"

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
