package main

import "loadbalancer/server"

func main() {
	server.RunConcurrentServers(3)
}
