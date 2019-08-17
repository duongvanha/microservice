package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/duongvanha/microservice/src/tools/obj"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Build struct {
	Dir  string `json:"dir"`
	Lang string `json:"lang"`
	Name string `json:"name"`
}

var buddyCmd = &cobra.Command{
	Use:   "buddy",
	Short: "gen config buddy",
	Run: func(cmd *cobra.Command, args []string) {
		var currentDir string
		if len(args) > 0 {
			currentDir = args[0]
		} else {
			currentDir, _ = os.Getwd()
		}

		var services []*Build

		err := filepath.Walk(currentDir, func(pathService string, info os.FileInfo, err error) error {
			if info.Name() == "build.json" {
				build, err := getBuild(pathService)
				if err != nil {
					fmt.Println("The file %w structure is incorrect", pathService)
					os.Exit(1)
				}
				services = append(services, build)
			}
			return nil
		})

		if err != nil {
			fmt.Println(fmt.Sprintf("Scan dir error %v", err))
		}

		err = BuildBuddyYml(services)

		if err != nil {
			fmt.Println(fmt.Sprintf("Error build yml file %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(buddyCmd)
}

func getBuild(path string) (*Build, error) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	build := Build{}

	err = json.Unmarshal([]byte(file), &build)
	if err != nil {
		return nil, err
	}

	return &build, nil
}

func BuildBuddyYml(services []*Build) error {
	var actions []obj.BuddyAction

	for _, service := range services {
		actions = append(actions, obj.BuddyAction{
			Pipeline:         service.Dir,
			TriggerMode:      "ON_EVERY_PUSH",
			RefName:          "master",
			RefType:          "BRANCH",
			TriggerCondition: "ALWAYS",
			Actions: []obj.Actions{
				{
					Action:           "Build Docker image",
					Type:             "DOCKERFILE",
					Login:            "${DOCKER_USERNAME}",
					Password:         "secure!+xmZnhe6DEcEPAlSioiycjtiz7E4mFqLbzs1Td60gKI=",
					Disabled:         false,
					DockerImageTag:   "latest,${BUDDY_EXECUTION_REVISION}",
					DockerfilePath:   fmt.Sprintf("src/docker/%s.dockerfile", service.Lang),
					ContextPath:      "src/services/" + service.Dir,
					Repository:       service.Name,
					TriggerCondition: "ALWAYS",
				},
			},
		})
	}

	output, err := yaml.Marshal(&actions)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	nameFile := "buddy.yml"

	writer, err := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	_, err = writer.Write(output)

	return err
}
