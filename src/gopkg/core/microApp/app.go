package micro_app

import (
	"flag"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/logger/zap/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
)

type MicroApp struct {
	name    string
	port    string
	version string
}

func init() {
	zapLogger, err := zap.NewLogger(zap.WithCallerSkip(2))
	if err != nil {
		panic(fmt.Sprintf("Cannot create logger with the following error: %s", err))
	}

	logger.DefaultLogger = zapLogger
}

func getMicroApp() MicroApp {
	app := MicroApp{}
	flag.StringVar(&app.name, "app-name", "go.haduong.api.movie", "App Name")
	flag.StringVar(&app.port, "http-host", ":3000", "HTTP listen host")
	flag.StringVar(&app.version, "version", "latest", "App Name")
	return app
}

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {

	microApp := getMicroApp()

	// set default options
	options := []micro.Option{
		micro.Version(microApp.version),
		micro.Name(microApp.name),
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
	service := NewService()
	microApp := getMicroApp()

	// prepend option
	options := []web.Option{
		web.MicroService(service),
		web.Address(microApp.port),
	}

	// append user options
	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
