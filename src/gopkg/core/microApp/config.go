package micro_app

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
)

var configSource config.Config

func GetConfig() (config.Config, error) {
	if configSource == nil {
		// Create new config
		conf, err := config.NewConfig()

		if err != nil {
			return nil, err
		}

		configSource = conf
	}

	return configSource, nil
}

func LoadSource(source source.Source) (err error) {
	if _, err = GetConfig(); err != nil {
		return
	}

	return configSource.Load(source)
}
