# Setup infra for learning

Begin project. I'm using terraform to create google cloud compute, ansible to create kubernetes Architecture. For learning.
But google support kubernetes Architecture ( gke ) -> using terraform create kubernetes Architecture and helm for update service

# Terraform

Terraform create Infrastructure by code ( gke ):

- copy file env/microservices.json (credential for google cloud)

- terraform plan for review infra, terraform apply for apply infra


# Tools

Write by go in src/tools

- Gen config cicd ( buddy.works ) for service `tools buddy`

- Gen value helm chart for service ( helm upgrade )
`cd {{service}} && helm upgrade --install -f values.yaml {{service} --set image.tag=latest ../../../infra/charts --namespace microservice`

- Gen gateway-route for all service ( kubectl apply )
`kubectl -n microservice apply -f infra/services/gateway.yaml`


## Notes issue

- Using same version openssl encrypt and decrypter
- Kubeadm ( setup hosted kubernetes cluster not using gke because begin learning 😭)
`https://github.com/kubernetes/kubeadm/issues/1390
Instead of using --apiserver-advertise-address=publicIp use --apiserver-cert-extra-sans=publicIp
Don't forget to replace the private ip for the public ip in your .kube/config if you use kubectl from remote.`

- Istio require memory and CPU
`As per "https://docs.giantswarm.io/guides/deploying-istio/"
Some of them, as pilot, have a large impact in terms of memory and CPU, so it is recommended to have around 8GB of memory and 4 CPUs free in your cluster. Obviously, all components have requested resources defined, so if you don’t have enough capacity you will see pods not starting.`

- Istio gateway
`https://istio.io/docs/tasks/traffic-management/ingress/ingress-control/#determining-the-ingress-ip-and-ports"
Some cloud provider support `

## Notes branch
- kube-adm-google
Branch stores the writing setup process infra on google cloud (not use gke) using ansible + terraform


## Notes document References

- http://docs.shippable.com/deploy/tutorial/create-kubeconfig-for-self-hosted-kubernetes-cluster/

- https://medium.com/@zhaohuabing/which-one-is-the-right-choice-for-the-ingress-gateway-of-your-service-mesh-21a280d4a29c

## Todo

- [ ] Grafana ( istio support )
- [ ] Prometheus
- [ ] Move to aws if have money
- [ ] Tools gen istio-gateway ( route for all services )
- [ ] Move cicd to jenkins or https://travis-ci.org ( buddy.works limit build 120/M )
- [ ] Setup elk ( log service )
- [ ] Setup CockroachDB ( test HA )
- [ ] Setup kafka
- [ ] More
