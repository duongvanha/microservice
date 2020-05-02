package micro_app

import (
	"flag"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
)

type MicroApp struct {
	name string
	port string
}

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {

	// set default options
	options := []micro.Option{
		micro.Version("latest"),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return micro.NewService(options...)
}

// NewService returns a web service for kubernetes
func NewWebApp(opts ...web.Option) web.Service {

	// create new service
	service := micro.NewService()
	app := MicroApp{}

	// flags for http
	flag.StringVar(&app.port, "http-host", ":3000", "HTTP listen host")
	flag.StringVar(&app.name, "app-name", "go.haduong.api.movie", "App Name")

	// prepend option
	options := []web.Option{
		web.MicroService(service),
		web.Name(app.name),
		web.Address(app.port),
	}

	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
