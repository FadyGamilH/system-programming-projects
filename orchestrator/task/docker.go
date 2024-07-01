package task

import (
	"context"
	"io"
	"log"
	"math"
	"orchestrator/messages"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Docker struct {
	Client *client.Client
	Config *DockerConfig
}

// this response will be used as the response to Stop() and Start() apis of our worker
type DockerResponse struct {
	ContainerId string
	Action      string // START / STOP
	Error       error
	Details     string
}

/*
RunContainer() represents the api of running container from the docker cli, so the steps are :
  - Pulling the image if its not pulled before
  - Create the container
  - Start the container
*/
func (d *Docker) RunContainer() DockerResponse {
	ctx := context.Background()

	img, err := d.Client.ImagePull(ctx, d.Config.Image, types.ImagePullOptions{})
	if err != nil {
		log.Println(messages.ErrorFailedToPullImage+" : {%v}", d.Config.Image)
		return DockerResponse{
			Error: err,
		}
	}
	// print the result of pulling the img to the terminal (as the service will be running on the terminal )
	io.Copy(os.Stdout, img)

	// get the configs ready for docker
	restartPolicyConfig := container.RestartPolicy{
		Name: container.RestartPolicyMode(d.Config.RestartPolicy),
	}

	resourcesConfig := container.Resources{
		Memory:   d.Config.Memory,
		NanoCPUs: int64(d.Config.CPU * math.Pow(10, 9)),
	}

	containerConfig := container.Config{
		Image:        d.Config.Image,
		Tty:          false,
		Env:          d.Config.Env,
		ExposedPorts: d.Config.ExposedPorts,
	}

	hostConfig := container.HostConfig{
		RestartPolicy:   restartPolicyConfig,
		Resources:       resourcesConfig,
		PublishAllPorts: true, // now docker will expose the available prot
	}

	resp, err := d.Client.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, d.Config.Name)
	if err != nil {
		log.Printf(messages.ErrorCreatingContainer+" image : {%v}, error : {%v}\n", d.Config.Image, err)
		return DockerResponse{Error: err}
	}

	err = d.Client.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		log.Printf(messages.ErrorStartingContainer+" container_id : {%v}, error : {%v}\n", resp.ID, err)
		return DockerResponse{Error: err}
	}

	return DockerResponse{}
}
