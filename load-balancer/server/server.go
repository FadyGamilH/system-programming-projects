package server

import (
	"fmt"
	"slices"
)

type loadBalancerServerList struct {
	ports []string
}

/*
This function creates new server lists with thier ports so our load balancer can server the trafic to them

  - Receives the number of servers
  - Create new port for each server and append it to the final ports of servers
  - All servers will be running on " 808X " port, so if client code sends numOfServers = 3, we will have []ports {8080, 8081, 8082}
*/
func NewLoadBalancerServerList(numOfServers int) *loadBalancerServerList {
	ports := []string{}
	for i := 0; i < numOfServers; i++ {
		ports = append(ports, fmt.Sprintf("808%d", i))
	}
	return &loadBalancerServerList{
		ports: ports,
	}
}

func (lbsl *loadBalancerServerList) PopServer() string {
	serverPort := lbsl.ports[0]
	lbsl.ports = slices.Delete(lbsl.ports, 0, 1)
	return serverPort
}
