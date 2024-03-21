package xguard

import (
	"github.com/ares0516/xguard/route"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func Test_Start(t *testing.T) {
	xg := NewXGuard()
	var item route.CommonRouteItem
	var itemList []route.CommonRouteItem
	item.Destination = "1.1.1.1"
	item.Mask = "255.255.255.255"
	item.NextHop = "192.168.31.1"
	item.InterfaceIndex = 11
	item.Metric = 26
	itemList = append(itemList, item)
	xg.AddRule(itemList)
	xg.Start()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

}
