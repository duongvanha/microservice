package cmd

import (
	"log"
	. "microservice/tools/obj"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestExecCmd(t1 *testing.T) {
	var t = []BuddyAction{
		{
			Pipeline:         "demo",
			TriggerMode:      "MANUAL",
			RefName:          "kafka/consumer",
			RefType:          "BRANCH",
			TriggerCondition: "ALWAYS",
			Actions: []Actions{
				{
					Action:           "Build Docker image",
					Type:             "DOCKERFILE",
					Login:            "${DOCKER_USERNAME}",
					Password:         "${DOCKER_PASSWORD}",
					Disabled:         false,
					DockerImageTag:   "latest,${BUDDY_EXECUTION_REVISION}",
					DockerfilePath:   "src/docker/go.dockerfile",
					ContextPath:      "src/services/simple",
					Repository:       "blademaster996/simple",
					TriggerCondition: "ALWAYS",
				},
			},
		},
	}

	output, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	nameFile := "demo.yml"

	writer, err := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		println(err)
	}

	_, err = writer.Write(output)

	if err != nil {
		println(err)
	}

}
