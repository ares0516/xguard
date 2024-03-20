package route

import "testing"

func Test_getRouteList(t *testing.T) {
	t.Parallel()
	_, err := getRouteList()
	if err != nil {
		t.Errorf("Failed to get route list: %v", err)
	}
}
