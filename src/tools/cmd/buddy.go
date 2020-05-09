package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"microservice/src/tools/helper"
	"microservice/src/tools/obj"
	"os"
)

var buddyCmd = &cobra.Command{
	Use:   "buddy",
	Short: "gen config buddy",
	Run: func(cmd *cobra.Command, args []string) {
		services, err := helper.ScannerService(args)
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

func BuildBuddyYml(services []*obj.Build) error {
	var actions []obj.BuddyAction

	for _, service := range services {
		actions = append(actions, obj.BuddyAction{
			Pipeline:         helper.GetSubPath(service.Dir, 1),
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
					ContextPath:      helper.GetSubPath(service.Dir, 3),
					Repository:       service.Name,
					TriggerCondition: "ALWAYS",
				},
			},
		})
	}

	var prefix []byte

	file, err := ioutil.ReadFile("base.yml")

	if err != nil {
		println("warning file base nil ", err)
	} else {
		prefix = file
	}

	output, err := yaml.Marshal(&actions)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	nameFile := "buddy.yml"

	writer, err := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	_, err = writer.Write(append(prefix, output...))

	return err
}
