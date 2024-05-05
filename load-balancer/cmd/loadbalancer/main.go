package main

import loadbalancer "loadbalancer/load-balancer"

func main() {
	loadbalancer.RunLoanBalancer(3)
}
