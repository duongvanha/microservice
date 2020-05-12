package main

import (
	"flag"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router/gin"
	microLogger "github.com/micro/go-micro/v2/logger"
	micro_app "microservice/src/gopkg/core/microApp"
	"os"
)

func main() {
	registry := micro_app.GetRegistry()

	services, err := registry.ListServices()

	if err != nil {
		microLogger.Fatal(err.Error())
	}

	var endpoints []*config.EndpointConfig

	for _, service := range services {
		for _, endpoint := range service.Endpoints {
			endpoints = append(endpoints, &config.EndpointConfig{
				Method:         "GET",
				Endpoint:       service.Nodes[0].Metadata["route"] + endpoint.Name,
				OutputEncoding: "no-op",
				Backend: []*config.Backend{{
					Method:                   "GET",
					Host:                     []string{"http://localhost:3000"},
					URLPattern:               endpoint.Name,
					SD:                       "static",
					HostSanitizationDisabled: true,
				}},
			})
		}
	}

	println(services)

	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "/Users/duongha/go/src/microservice/src/services/apiGateway/config.json", "Path to the configuration filename")
	flag.Parse()

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		microLogger.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	serviceConfig.Endpoints = endpoints

	logger, _ := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")

	routerFactory := gin.DefaultFactory(proxy.DefaultFactory(logger), logger)

	routerFactory.New().Run(serviceConfig)
}

// customProxyFactory adds a logging middleware wrapping the internal factory
type customProxyFactory struct {
	logger  logging.Logger
	factory proxy.Factory
}

// New implements the Factory interface
func (cf customProxyFactory) New(cfg *config.EndpointConfig) (p proxy.Proxy, err error) {
	p, err = cf.factory.New(cfg)
	if err == nil {
		p = proxy.NewLoggingMiddleware(cf.logger, cfg.Endpoint)(p)
	}
	return
}
