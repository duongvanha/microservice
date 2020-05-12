module microservice/src

go 1.13

require (
	github.com/devopsfaith/krakend v1.1.1
	github.com/devopsfaith/krakend-etcd v0.0.0-20190425091451-d989a26508d7
	github.com/gin-gonic/gin v1.1.5-0.20170702092826-d459835d2b07
	github.com/golang/protobuf v1.3.5
	github.com/micro/go-micro/v2 v2.5.0
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.5.0
	github.com/micro/go-plugins/logger/zap/v2 v2.5.0
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.5.0
	github.com/spf13/cobra v0.0.5
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/ugorji/go v0.0.0-20180112141927-9831f2c3ac10 => github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8
