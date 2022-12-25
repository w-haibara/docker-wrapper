package main

import (
	"fmt"
	"log"

	"github.com/w-haibara/docker-wrapper/docker"
)

func main() {
	noCache := true
	cmd := docker.DockerBuildCmd(docker.DockerBuildOption{
		NoCache: &noCache,
		Tag:     []string{"name1:tag1", "name2:tag2"},
	}, []string{"."})

	log.Println("command:", cmd)

	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}
}
