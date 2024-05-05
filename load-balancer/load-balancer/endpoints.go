package loadbalancer

import "net/url"

type endpoints struct {
	urls []url.URL
}

/*
RoundRobinShufelling() is a receiver function on endpoints type

  - perform round robin shuffeling by taking the first url and put it at the end of the list so the next url is the 2nd url and so on
*/
func (e *endpoints) RoundRobinShufelling() {
	firstUrl := e.urls[0]
	e.urls = e.urls[1:]
	e.urls = append(e.urls, firstUrl)
}

func (e *endpoints) CurrentEndpoint() *url.URL {
	return &e.urls[0]
}
