package task

import (
	"github.com/docker/docker/client"
)

type Docker struct {
	Client *client.Client
	Config *Config
}
