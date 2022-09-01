package main

import (
	"fmt"
	"sync"
)

// Usage:
// pool.Add(portal)
// portal := pool.Pick()
type PortalPool struct {
	m      []*Portal
	rptr   int
	wptr   int
	cnt    int
	mlen   int
	maxlen int

	mu sync.Mutex
}

func NewPortalPool(initsize int, maxsize int) *PortalPool {
	m := make([]*Portal, initsize)
	return &PortalPool{
		m:      m,
		rptr:   0,
		wptr:   0,
		cnt:    0,
		mlen:   initsize,
		maxlen: maxsize,
	}
}
func (p *PortalPool) Cnt() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.cnt
}
func (p *PortalPool) Add(portal *Portal) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cnt >= p.maxlen {
		p.m[p.wptr] = portal
		p.rptr = (p.rptr + 1) % p.mlen
		p.wptr = (p.wptr + 1) % p.mlen
	} else if p.mlen == p.cnt {
		p.m = append(p.m, portal)
		p.mlen++
		p.cnt++
	} else {
		p.m[p.wptr] = portal
		p.wptr = (p.wptr + 1) % p.mlen
		p.cnt++
	}
}
func (p *PortalPool) Pick() (portal *Portal) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cnt == 0 {
		// portal = nil
		return nil
	} else {
		portal = p.m[p.rptr]
		p.m[p.rptr] = nil
		p.rptr = (p.rptr + 1) % p.mlen
		p.cnt--
	}
	return
}

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
