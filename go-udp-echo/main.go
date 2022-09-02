package main

import (
	"log"
	"net"
)

func main() {
	la, _ := net.ResolveUDPAddr(`udp`, `:9999`)
	lc, _ := net.ListenUDP(`udp`, la)
	buf := make([]byte, 2048)
	for {
		l, addr, err := lc.ReadFrom(buf)
		log.Println(l, addr)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		lc.WriteTo(buf[:l], addr)
	}
}
