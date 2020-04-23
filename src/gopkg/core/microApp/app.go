package micro_app

import (
	"flag"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/memory"
	"github.com/micro/go-micro/v2/web"
)

type MicroApp struct {
	name string
	port string
}

var localRegistry registry.Registry

func getRegistry() registry.Registry {
	if localRegistry == nil {
		localRegistry = memory.NewRegistry()
	}
	return localRegistry
}

func (app MicroApp) Init() error {
	return nil
}

func (app MicroApp) Run() error {
	return nil
}

func (app MicroApp) OnClose(func() error) {

}

func NewHTTPApp() web.Service {

	//file, err := afero.ReadFile(v.fs, filename)

	app := MicroApp{}

	// flags for http
	flag.StringVar(&app.port, "http-host", ":3000", "HTTP listen host")
	flag.StringVar(&app.name, "app-name", "shippy.service.api", "App Name")

	return web.NewService(
		web.Name(app.name),
		web.Address(app.port),
	)

}
