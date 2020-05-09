package helper

import (
	"fmt"
	"microservice/src/tools/obj"
	"os"
	"path/filepath"
	"strings"
)

func ScannerService(args []string) (services []*obj.Build, err error) {
	var currentDir string
	if len(args) > 0 {
		currentDir = args[0]
	} else {
		currentDir, _ = os.Getwd()
	}

	err = filepath.Walk(currentDir, func(pathService string, info os.FileInfo, err error) error {
		if info.Name() == "build.json" {
			build, err := obj.GetBuild(pathService)
			if err != nil {
				fmt.Println("The file %w structure is incorrect", pathService)
				return nil
			}

			if build.Dir == "" {
				build.Dir = filepath.Dir(pathService)
			}
			services = append(services, build)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return services, nil
}

func GetSubPath(dir string, level uint64) string {
	a := strings.Split(dir, "/")

	position := uint64(len(a)) - level

	a = a[position:]

	return strings.Join(a, "/")
}
