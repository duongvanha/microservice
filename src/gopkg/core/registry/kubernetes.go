package registry

import (
	"github.com/micro/go-micro/v2/registry"
)

type Kubernetes struct {
}

func (k Kubernetes) Init(option ...registry.Option) error {
	panic("implement me")
}

func (k Kubernetes) Options() registry.Options {
	panic("implement me")
}

func (k Kubernetes) Register(service *registry.Service, option ...registry.RegisterOption) error {
	panic("implement me")
}

func (k Kubernetes) Deregister(service *registry.Service, option ...registry.DeregisterOption) error {
	panic("implement me")
}

func (k Kubernetes) GetService(s string, option ...registry.GetOption) ([]*registry.Service, error) {
	panic("implement me")
}

func (k Kubernetes) ListServices(option ...registry.ListOption) ([]*registry.Service, error) {
	panic("implement me")
}

func (k Kubernetes) Watch(option ...registry.WatchOption) (registry.Watcher, error) {
	panic("implement me")
}

func (k Kubernetes) String() string {
	panic("implement me")
}
