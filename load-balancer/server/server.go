package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"slices"
	"sync"

	"github.com/gin-gonic/gin"
)

type loadBalancerServerList struct {
	ports []string
}

// Global map to store server instances by port
var serversMap = make(map[string]*http.Server)

/*
This function creates new server lists with thier ports so our load balancer can server the trafic to them

  - Receives the number of servers
  - Create new port for each server and append it to the final ports of servers
  - All servers will be running on " 808X " port, so if client code sends numOfServers = 3, we will have []ports {8080, 8081, 8082}
*/
func NewLoadBalancerServerList(numOfServers int16) *loadBalancerServerList {
	ports := []string{}
	for i := 0; i < int(numOfServers); i++ {
		ports = append(ports, fmt.Sprintf("808%d", i))
	}
	return &loadBalancerServerList{
		ports: ports,
	}
}

/*
This function removes the first server from the serverList and return its port to the caller

  - If current state of the servers is []ports{8080, 8081} after calling it , we will receive 8080 and the state will be []ports{8081}
*/
func (lbsl *loadBalancerServerList) PopServer() string {
	serverPort := lbsl.ports[0]
	lbsl.ports = slices.Delete(lbsl.ports, 0, 1)
	return serverPort
}

func RunConcurrentServers(numOfConcurrentServers int16) error {
	servers := NewLoadBalancerServerList(numOfConcurrentServers)

	var wg sync.WaitGroup
	wg.Add(int(numOfConcurrentServers))
	defer wg.Wait()

	for i := 0; i < int(numOfConcurrentServers); i++ {
		go StartServer(servers, &wg)
	}

	return nil
}

func StartServer(servers *loadBalancerServerList, wg *sync.WaitGroup) {
	defer wg.Done()

	var mu sync.RWMutex
	var port string
	mu.Lock()
	{
		port = servers.PopServer()
	}
	mu.Unlock()

	router := gin.Default()
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Store the server instance in the map
	serversMap[port] = httpServer

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"server": port,
		})
	})

	router.GET("/shutdown", func(c *gin.Context) {
		// Retrieve the server instance from the map
		srv, ok := serversMap[port]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
			return
		}
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Server Shutdown Failed:%+v", err)
		}
	})

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
