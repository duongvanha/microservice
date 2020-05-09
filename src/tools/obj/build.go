package obj

import (
	"encoding/json"
	"io/ioutil"
)

type Build struct {
	Dir  string `json:"dir"`
	Lang string `json:"lang"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetBuild(path string) (*Build, error) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	build := Build{}

	err = json.Unmarshal(file, &build)
	if err != nil {
		return nil, err
	}

	if build.Dir == "" {

	}

	return &build, nil
}
