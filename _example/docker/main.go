package main

import (
	"fmt"
	"log"

	"github.com/w-haibara/docker-wrapper/docker"
)

func main() {
	cmd := docker.DockerCmd(docker.DockerOption{}, []string{})
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
