package main

import loadbalancer "loadbalancer/loadbalancer"

func main() {
	loadbalancer.RunLoanBalancer(3)
}
