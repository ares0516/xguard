package route

type CommonRouteTable struct {
	Items         []CommonRouteItem
	DefaultMetric int
}

type CommonRouteItem struct {
	Destination    string
	Mask           string
	NextHop        string
	InterfaceIndex int
	Metric         int
}

// Contains checks if the route table contains the given route item
func (cr *CommonRouteTable) Contains(item CommonRouteItem) bool {
	for _, i := range cr.Items {
		if i.Destination == item.Destination && i.Mask == item.Mask && i.NextHop == item.NextHop && i.InterfaceIndex == item.InterfaceIndex {
			return true
		}
	}
	return false
}

func GetRouteTable() (*CommonRouteTable, error) {
	routeTable, err := getIPv4RouteTable()
	if err != nil {
		return nil, err
	}
	return routeTable, nil
}

func SetRouteList(routeTable *CommonRouteTable) error {
	return setRouteList(routeTable)
}
