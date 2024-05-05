package loadbalancer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var (
	BaseUrl          = "http://localhost:808"
	loadBalancerPort = "3000"
)

type loadBalancer struct {
	reverseProxy     *httputil.ReverseProxy
	serversEndpoints endpoints
}

/*
Factory method used to return new loadbalancer instance
  - receives the number of servers to be balanced
  - create endpoints for all servers, append them to the laodbalancer instance and return it
*/
func newLoadBalancer(numOfServers int16) *loadBalancer {
	endpoints := endpoints{
		urls: []url.URL{},
	}

	for i := 0; i < int(numOfServers); i++ {
		url, err := url.Parse(fmt.Sprintf("%s%d/", BaseUrl, i))
		if err != nil {
			log.Fatal(err)
		}
		endpoints.urls = append(endpoints.urls, *url)
	}
	return &loadBalancer{
		serversEndpoints: endpoints,
	}
}

type ServerResponse struct {
	Server string `json:"server"`
}

/*
RunLoanBalancer(numOfServers) start the loadbalancer http server

  - define http server for our loadbalancer
  - check the current server if its up and ready to receive the trafik, route the request to it, otherwise we roundrobin and check health of the next one, until we find the healthy server
  - roundrobin the servers
*/
func RunLoanBalancer(numOfServers int16) {
	router := gin.Default()

	lb := newLoadBalancer(numOfServers)

	router.GET("/loadbalancer", func(c *gin.Context) {
		for !isServerUp(*lb.serversEndpoints.CurrentEndpoint()) {
			// shuffel so we check the health of the next one
			lb.serversEndpoints.RoundRobinShufelling()
		}
		currentServerUrl := lb.serversEndpoints.CurrentEndpoint()
		lb.reverseProxy = httputil.NewSingleHostReverseProxy(currentServerUrl)
		// once we here, so the server is up and running, and we cna redirect the request to it, then we need to shuffel the servers to forward the next request to the next server
		lb.serversEndpoints.RoundRobinShufelling()
		// lb.reverseProxy.ServeHTTP(c.Writer, c.Request)
		// Manually construct the request to forward
		req, err := http.NewRequest("GET", currentServerUrl.String(), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Use the http.Client to send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		defer resp.Body.Close()

		// Forward the response back to the client
		// Read the response body
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Parse the JSON response body into the struct
		var bodyStruct ServerResponse
		err = json.Unmarshal(bodyBytes, &bodyStruct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Send the struct as part of the response
		c.JSON(http.StatusOK, gin.H{"resp": bodyStruct})
	})

	router.Run("localhost:" + loadBalancerPort)
}

/*
isServerUp(url) check if the server is up and running, and its response is 200
*/
func isServerUp(serverUrl url.URL) bool {
	response, err := http.Get(serverUrl.String())
	if err != nil {
		return false
	}
	return response.StatusCode == http.StatusOK
}
