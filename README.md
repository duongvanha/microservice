# Setup infra for learning

# Terraform

Terraform create Infrastructure by code ( gke ):

- copy file env/microservices.json (credential for google cloud)

- terraform plan for review infra, terraform apply for apply infra


# Tools

Write by go in src/tools

- Gen config cicd ( buddy.works ) for service `tools buddy`

- Gen value helm chart for service ( helm upgrade )
```shell script
cd {{service}} && helm upgrade --install -f values.yaml {{service-name}} ../../../infra/microservices`
```

-- Gen proto
```shell script
cd $GOPATH/src/microservice/src/proto
protoc --proto_path=$GOPATH/src:. --micro_out=../gopkg/ --go_out=../gopkg/ services/*.proto 
protoc --proto_path=$GOPATH/src:. --micro_out=../../.. --go_out=../../.. models/*.proto
```

-- build docker
```shell script
cd ~/go/src/microservice/src
docker build . -f docker/go.dockerfile --no-cache --build-arg dir=${dir_service} -t ${name_service}:${tag}
#docker build . -f docker/go.dockerfile --no-cache --build-arg dir=services/buildMovie -t blademaster996/go.haduong.api.movie:1
```

-- tunnel deployment
```shell script
# run script command
micro tunnel --server=${ip}:${port}
#run service local with env MICRO_PROXY=go.micro.tunnel
MICRO_PROXY=go.micro.tunnel go run main.go
```

- Apply gateway-route for all service ( kubectl apply )
```shell script
kubectl -n microservice apply -f infra/services/gateway.yaml
```

## Notes issue

- Using the same version OpenSSL encrypt and decrypter
- kubeadm ( setup hosted Kubernetes cluster not using gke because begin learning 😭)
`https://github.com/kubernetes/kubeadm/issues/1390
Instead of using --apiserver-advertise-address=publicIp use --apiserver-cert-extra-sans=publicIp
Don't forget to replace the private IP for the public IP in your .kube/config if you use kubectl from remote.`

- Istio require memory and CPU
`As per "https://docs.giantswarm.io/guides/deploying-istio/"
Some of them, as the pilot, have a large impact in terms of memory and CPU, so it is recommended to have around 8GB of memory and 4 CPUs free in your cluster. Obviously, all components have requested resources defined, so if you don’t have enough capacity you will see pods not starting.`

- Istio gateway
`https://istio.io/docs/tasks/traffic-management/ingress/ingress-control/#determining-the-ingress-ip-and-ports"
Some cloud provider support `

- Solution for Development process using native kubernetes service discovery ( replace consul )

## Notes branch
- kube-adm-google
Branch stores the writing setup process infra on google cloud (not use gke) using ansible + terraform


## Notes document References

- http://docs.shippable.com/deploy/tutorial/create-kubeconfig-for-self-hosted-kubernetes-cluster/

- https://medium.com/@zhaohuabing/which-one-is-the-right-choice-for-the-ingress-gateway-of-your-service-mesh-21a280d4a29c

- https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

- https://stackoverflow.com/questions/54698875/gcloud-kubernetes-cluster-with-1-insufficient-cpu-error

- https://www.digitalocean.com/community/tutorials/how-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes

- https://mherman.org/blog/logging-in-kubernetes-with-elasticsearch-Kibana-fluentd/

- https://www.cockroachlabs.com/docs/stable/orchestrate-cockroachdb-with-kubernetes.html#hosted-gke

- https://github.com/micro/development
## Todo infra
- [x] istio gateway
- [x] Grafana ( istio support ) -> remove isto, setup manual
- [x] Prometheus
- [x] Setup elk ( log service ) -> improve log error stacktrace
- [x] Setup CockroachDB ( test HA )
- [ ] Setup kafka
- [x] Service discovery for developments
- [ ] Tools gen istio-gateway ( route for all services ) -> remove istio using api gateway
- [ ] Move cicd to jenkins or https://travis-ci.org ( buddy.works limit build 120/M )
- [ ] Move to aws if have money
- [ ] More
------------------------
## Todo update full application
- [ ] Build consumer scanner movie
- [ ] Build consumer update movie
- [ ] Build graphl api
- [ ] Build rest api
- [ ] Build mobile app with flutter
- [ ] Build dashboard with angular
