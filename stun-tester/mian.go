package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var c *net.UDPConn
var peerAddr *net.UDPAddr

var getPool *PortalPool
var putPool *PortalPool

var m map[string]*Portal
var mu sync.Mutex

func main() {
	// addr := "localhost:9999"
	// Server(addr)

	addr := ":10000"
	Client(addr)
}

func ActiveClientPortal() {
	if peerAddr == nil {
		return
	}

	p := getPool.Pick()
	if p == nil {
		return
	}

	fmt.Println("ActiveClientPortal", p, getPool, putPool, peerAddr, c)

	c.WriteToUDP([]byte(p.LocalAddr), peerAddr)
	mu.Lock()
	m[p.LocalAddr] = p
	mu.Unlock()

	fmt.Println("ActiveClientPortal", p, getPool, putPool)

	// for i := putPool.mlen - putPool.Cnt(); i > 0; i-- {
	// 	ActiveClientPortal()
	// }
	if putPool.mlen > putPool.Cnt() {
		go ActiveClientPortal()
	}
}

func Client(listenAddr string) {
	stopFlag := false
	pc := NewPortalClient(listenAddr)
	for i := 0; i < 5; i++ {
		go pc.NewPortal()
	}
	go debug(pc)
	getPool = pc.Pool
	putPool = pc.Mux.Pool
	pc.Mux.RecvConnCallBack = ActiveClientPortal

	var err error
	c, err = net.ListenUDP("udp", nil)
	if err != nil {
		log.Println(err)
		return
	}
	localAddr, err := GetAddr(c)
	if err != nil {
		log.Println(err)
		return
	}

	go func(host string, path string, data string) { // not tested, should be work. connect between
		for !stopFlag {
			res := PostAddr(host, path, data, 10)
			if res != nil && len(res) == 2 {
				for _, s := range res {
					udpaddr, err := net.ResolveUDPAddr("udp", s)
					if err != nil {
						log.Println(err)
					}
					c.WriteToUDP([]byte{}, udpaddr)
					if s != data {
						peerAddr = udpaddr
					}
				}
				time.Sleep(time.Second * 3)
			}
		}
	}("http://127.0.233.0:8080", "test", localAddr)

	m = make(map[string]*Portal)
	for peerAddr == nil {
		time.Sleep(time.Second)
	}
	ActiveClientPortal()

	buf := make([]byte, 2048)
	for {
		l, raddr, err := c.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		if l == 0 {
			continue
		}
		msg := string(buf[:l])
		msgs := strings.Split(msg, "\n")
		_, err = net.ResolveUDPAddr("udp", msgs[1])
		if err != nil {
			log.Println(err)
			continue
		}
		_, err = net.ResolveUDPAddr("udp", msgs[0])
		if err != nil {
			log.Println(err)
			continue
		}

		mu.Lock()
		p := m[msgs[0]]
		if p != nil {
			p.Set(&msgs[1], nil, pc.Mux) // set peer only, local addr will set by mux
			p.Mux.Pool.Add(p)
			delete(m, msgs[0])
		}
		mu.Unlock()

		_ = raddr
	}

}

func Server(forwardAddr string) {
	stopFlag := false
	ps := NewPortalServer(forwardAddr)
	for i := 0; i < 5; i++ {
		go ps.NewPortal()
	}

	var err error
	c, err = net.ListenUDP("udp", nil)
	if err != nil {
		log.Println(err)
		return
	}
	localAddr, err := GetAddr(c)
	if err != nil {
		log.Println(err)
		return
	}
	go func(host string, path string, data string) { // not tested, should be work. connect between
		for !stopFlag {
			res := PostAddr(host, path, data, 10)
			if res != nil {
				for _, s := range res {
					udpaddr, err := net.ResolveUDPAddr("udp", s)
					if err != nil {
						log.Println(err)
					}
					c.WriteToUDP([]byte{}, udpaddr)
				}
				time.Sleep(time.Second * 3)
			}
		}
	}("http://127.0.233.0:8080", "test", localAddr)

	// loop
	// handle the recv message
	buf := make([]byte, 2048)
	for {
		l, raddr, err := c.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		if l == 0 {
			continue
		}
		msg := string(buf[:l])
		_, err = net.ResolveUDPAddr("udp", msg)
		if err != nil {
			log.Println(err)
			continue
		}
		p := ps.ActivePortal(&msg, ps.LocalAddr, nil)

		res := msg + "\n" + p.LocalAddr
		c.WriteToUDP([]byte(res), raddr)
		c.WriteToUDP([]byte(res), raddr)
		c.WriteToUDP([]byte(res), raddr)
	}
}

func debug(pc *PortalClient) {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println(pc.Mux.Pool)
		fmt.Println(pc.Pool)
		fmt.Println("===============")
	}
}
