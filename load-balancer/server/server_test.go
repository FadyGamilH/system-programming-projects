package server

import "testing"

func Test_ServersFactory(t *testing.T) {
	servers := NewLoadBalancerServerList(2)
	expected := []string{"8080", "8081"}
	actual := servers.ports
	if len(expected) != len(actual) {
		t.Errorf("got : %v, expected : %v", actual, expected)
	}
	for idx, port := range actual {
		if port != expected[idx] {
			t.Errorf("got : %v, expected : %v", port, expected[idx])
		}
	}
}
