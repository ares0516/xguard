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

func GetRouteList() (CommonRouteTable, error) {
	var routeTable CommonRouteTable
	routeList, err := getRouteList()
	if err != nil {
		return CommonRouteTable{}, err
	}
	routeTable.Items = append(routeTable.Items, routeList...)
	return routeTable, nil
}

func SetRouteList(routeTable *CommonRouteTable) error {
	return setRouteList(routeTable)
}
