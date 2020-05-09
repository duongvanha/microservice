package cmd

import (
	"microservice/src/tools/helper"
	"path/filepath"
	"testing"
)

func TestExecCmd(t1 *testing.T) {
	dir := filepath.Dir("/Users/duongha/go/src/microservice/src/services/movieApi/build.json")

	println(helper.GetSubPath(dir, 3))

}
