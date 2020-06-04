module github.com/giantswarm/go-webhook-app

go 1.14

require (
	github.com/coreos/prometheus-operator v0.39.0 // indirect
	github.com/oklog/run v1.1.0
	github.com/prometheus/client_golang v1.6.0
	github.com/sirupsen/logrus v1.6.0
	github.com/slok/go-http-metrics v0.6.1 // indirect
	github.com/slok/kubewebhook v0.10.0 // indirect
	github.com/stretchr/testify v1.5.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
)
