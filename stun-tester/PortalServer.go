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
	p := NewPortal("udp")
	s.Pool.Add(p)
}

func (s *PortalServer) ActivePortal(paddr *string) {
	p := s.Pool.Pick()
	if p == nil {
		return
	}
	p.Set(paddr, s.LocalAddr, nil)
	if s.Pool.cnt < s.Pool.mlen {
		go s.NewPortal()
	}
}
