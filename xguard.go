package xguard

import (
	"net"
	"sync"
	"time"
)

type Rule struct {
	subnet net.IPNet
	gw     net.IP
	ifName string
	ifIdx  uint32
}

type XGuard struct {
	mu        sync.Mutex
	guardList []Rule
}

func NewXGuard() *XGuard {
	return &XGuard{}
}

func (x *XGuard) AddRule(subnet string, gw string, ifName string, ifIdx uint32) {
	x.mu.Lock()
	defer x.mu.Unlock()

	_, subnet, _ = net.ParseCIDR(subnet)
	gw := net.ParseIP(gw)
	x.guardList = append(x.guardList, Rule{subnet: *subnet, gw: gw, ifName: ifName, ifIdx: ifIdx})
}

func (x *XGuard) DelRule(subnet string) {
	x.mu.Lock()
	defer x.mu.Unlock()

	_, subnet, _ = net.ParseCIDR(subnet)
	for i, rule := range x.guardList {
		if rule.subnet.String() == subnet.String() {
			x.guardList = append(x.guardList[:i], x.guardList[i+1:]...)
			break
		}
	}
}

func (x *XGuard) CleanRule() {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.guardList = nil
}

func (x *XGuard) Start() {
	go func() {
		for {
			time.Sleep(5 * time.Second)

		}
	}()
}

func (x *XGuard) Stop() {
}

func (x *XGuard) Sync() {
}
