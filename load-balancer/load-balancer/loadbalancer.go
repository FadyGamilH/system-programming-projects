package loadbalancer

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

var (
	BaseUrl = "http://localhost:808"
)

type loadBalancer struct {
	reverseProxy     httputil.ReverseProxy
	serversEndpoints endpoints
}

func New(numOfServers int16) (*loadBalancer, error) {
	endpoints := endpoints{
		urls: []url.URL{},
	}

	for i := 0; i < int(numOfServers); i++ {
		url, err := url.Parse(fmt.Sprintf("%s%d", BaseUrl, i))
		if err != nil {
			return nil, err
		}
		endpoints.urls = append(endpoints.urls, *url)
	}
	return &loadBalancer{
		serversEndpoints: endpoints,
	}, nil
}
