package micro_app

import (
	"encoding/json"
	"io/ioutil"
)

type Build struct {
	Dir  string `json:"dir"`
	Lang string `json:"lang"`
	Name string `json:"name"`
	Type string `json:"name"`
	Port int64  `json:"name"`
}

func GetBuild(path string) (*Build, error) {
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
