package main

type PortalServer struct {
	LocalAddr *string
	Pool      *PortalPool
}

func NewPortalServer(addr string) *PortalServer {
	pool := NewPortalPool(5, 10)
	return &PortalServer{
		LocalAddr: &addr,
		Pool:      pool,
	}
}

func (s *PortalServer) NewPortal() {
	p := NewPortal("udp6 not supported")
	s.Pool.Add(p)
}

func (s *PortalServer) ActivePortal(paddr *string) {
	p := s.Pool.Pick()
	p.Set(paddr, s.LocalAddr, nil)
}
