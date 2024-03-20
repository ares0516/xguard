package xguard

import (
	"fmt"
	"gihub.com/ares0516/xguard/route"
	"sync"
	"time"
)

type XGuard struct {
	mu        sync.Mutex
	lastList  []route.CommonRouteItem // 上一次的路由表
	guardList []route.CommonRouteItem // 需要保护的路由表
}

func NewXGuard() *XGuard {
	return &XGuard{}
}

func (x *XGuard) AddRule(itemList []route.CommonRouteItem) {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.guardList = append(x.guardList, itemList...)
}

func (x *XGuard) DelRule(subnet string) {
	x.mu.Lock()
	defer x.mu.Unlock()

}

func (x *XGuard) CleanRule() {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.guardList = nil
}

func (x *XGuard) Start() {
	go func() {
		for {
			tmplist, err := route.GetRouteList()
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			}
			// 获取存在于guardList但不存在于tmplist的路由
			var routeTable route.CommonRouteTable
			for _, r := range x.guardList {
				found := false
				if tmplist.Contains(r) {
					found = true // 当前被保护条目存在，不处理
				}
				if !found { // 当前被保护条目不存在，添加路由
					fmt.Println("Add route:", r)
					routeTable.Items = append(routeTable.Items, r)
				}
			}
			route.SetRouteList(&routeTable)
			time.Sleep(5 * time.Second)
		}
	}()
}

func (x *XGuard) Stop() {
}

func (x *XGuard) Sync() {
}
