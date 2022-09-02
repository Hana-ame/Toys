package main

import "net"

type Multiplexer struct {
	Conn *net.UDPConn // 占用的port
}

func (m *Multiplexer) Remove(i interface{}) {

}

type PortalClient struct {
	Pool *PortalPool
	mux  *Multiplexer
}

func NewPortalClient(addr string) *PortalClient {
	// pool := NewPortalPool(5, 10)
	return &PortalClient{
		// LocalAddr: &addr,
		// Pool:      pool,
	}
}

func (s *PortalClient) NewPortal() {
	p := NewPortal("udp")
	s.Pool.Add(p)
}

func (s *PortalClient) ActivePortal(paddr *string) {
	// p := s.Pool.Pick()
	// p.Set(paddr, s.LocalAddr, nil)
}
