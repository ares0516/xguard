//go:build windows

package route

import (
	"fmt"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/StackExchange/wmi"
	SP "golang.org/x/text/encoding/simplifiedchinese"
)

type Win32_IP4RouteTable struct {
	Age            int
	Destination    string
	Mask           string
	NextHop        string
	InterfaceIndex int
	Metric1        int
	Metric2        int    // x
	Metric3        int    // x
	Metric4        int    // x
	Caption        string // x
	Description    string // x
	Name           string // x
	Status         string // x
	Information    string // x
}

// eg. route ADD 157.0.0.0 MASK 255.0.0.0  157.55.80.1 METRIC 3 IF 2

func setRouteList(routeTable *CommonRouteTable) error {
	baseMetric := routeTable.DefaultMetric
	fmt.Println("baseMetric:", baseMetric)
	for _, item := range routeTable.Items {
		metricOffset := item.Metric - baseMetric
		fmt.Println("metricOffset:", metricOffset)
		cmd := exec.Command("route", "add", item.Destination, "mask", item.Mask, item.NextHop, "metric", strconv.Itoa(metricOffset), "if", strconv.Itoa(item.InterfaceIndex))
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, err := cmd.CombinedOutput()
		if err != nil {
			//return err
		}
		str, _ := SP.GBK.NewDecoder().String(string(out))
		fmt.Println("exec succ ! {%s}[%s]", cmd.String(), str)
	}
	return nil
}

func getIPv4RouteTable() (*CommonRouteTable, error) {
	routeTable := new(CommonRouteTable)
	var win32Routes []Win32_IP4RouteTable
	query := "SELECT * FROM Win32_IP4RouteTable"
	err := wmi.Query(query, &win32Routes)
	if err != nil {
		fmt.Println("Failed to query Win32_IP4RouteTable:", err)
		return nil, err
	}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> 0")
	baseMetric := 500
	for _, route := range win32Routes {
		//fmt.Printf("%s/%s  [%s]\n", route.Destination, route.Mask, route.Status)
		//fmt.Printf("Metrics: %d %d %d %d\n", route.Metric1, route.Metric2, route.Metric3, route.Metric4)
		//fmt.Printf("NextHop: %s\n", route.NextHop)
		//fmt.Printf("Age: %d\n", route.Age)
		//fmt.Printf("InterfaceIndex: %d\n", route.InterfaceIndex)

		//fmt.Printf(" %s %s %s %d %d\n", route.Destination, route.Mask, route.NextHop, route.InterfaceIndex, route.Metric1)
		var routeItem CommonRouteItem

		routeItem.Destination = route.Destination
		routeItem.Mask = route.Mask
		routeItem.NextHop = route.NextHop
		routeItem.InterfaceIndex = route.InterfaceIndex
		routeItem.Metric = route.Metric1
		if route.Metric1 < baseMetric {
			baseMetric = route.Metric1
		}
		routeTable.Items = append(routeTable.Items, routeItem)

	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> 1")
	routeTable.DefaultMetric = baseMetric
	fmt.Println("routeTable.DefaultMetric:", routeTable.DefaultMetric)
	return routeTable, nil
}
