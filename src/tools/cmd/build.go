package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Build struct {
	Dir  string `json:"dir"`
	Lang string `json:"lang"`
	Name string `json:"name"`
}

var rootCmd = &cobra.Command{
	Use: "build",
	Run: func(cmd *cobra.Command, args []string) {
		var currentDir string
		if len(args) > 0 {
			currentDir = args[0]
		} else {
			currentDir, _ = os.Getwd()
		}

		println(currentDir)

		err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
			if info.Name() == "build.json" {
				ExecService(path)
			}
			return nil
		})

		if err != nil {
			println(fmt.Sprintf("Error %v", err))
		}
	},
}

func ExecService(pathService string) {
	file, err := ioutil.ReadFile(pathService)

	if err != nil {
		fmt.Println("The file %w is not found", pathService)
		os.Exit(1)
	}

	buildId := os.Getenv("TRAVIS_BUILD_NUMBER")

	if buildId == "" {
		buildId = "latest"
	}

	build := Build{}

	err = json.Unmarshal([]byte(file), &build)
	if err != nil {
		fmt.Println("The file %w structure is incorrect", pathService)
		os.Exit(1)
	}

	dockerFile := path.Join(pathService, fmt.Sprintf("../../../docker/%s.dockerfile", build.Lang))

	fmt.Println("Build Docker -->")
	err = ExecCmd(pathService, "docker", "build", ".", "-t", fmt.Sprintf("%s:%s", build.Name, buildId), fmt.Sprintf("-f=%s", dockerFile))
	if err != nil {
		fmt.Println("Error build docker %w", err)
		os.Exit(1)
	}
	fmt.Println("Push Docker -->")
	err = ExecCmd(pathService, "docker", "push", fmt.Sprintf("%s:%s", build.Name, buildId))
	if err != nil {
		fmt.Println("Push image docker %w", err)
		os.Exit(1)
	}

}

func ExecCmd(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = path.Dir(dir)

	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
