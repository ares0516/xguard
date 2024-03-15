package xguard

import (
	"sync"
	"sync/atomic"
)

type Guard struct {
	status atomic.Value
	mu     sync.Mutex
}

func NewGuard() *Guard {
	g := &Guard{}
	g.status.Store("stopped")
	return g
}

func (g *Guard) Lock() {
	g.mu.Lock()
}

func (g *Guard) Unlock() {
	g.mu.Unlock()
}

func (g *Guard) Start() {
	g.status.Store("running")
}

func (g *Guard) Stop() {
	g.status.Store("stopped")
}

func (g *Guard) GetStatus() string {
	return g.status.Load().(string)
}
