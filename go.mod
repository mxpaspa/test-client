module github.com/test-client

go 1.16

require (
	github.com/operator-framework/api v0.10.5
	github.com/redhat-marketplace/redhat-marketplace-operator/v2 v2.0.0-20210521154335-fbc63727152a
	k8s.io/apimachinery v0.22.0
)

replace k8s.io/client-go => k8s.io/client-go v0.19.4

replace k8s.io/api => k8s.io/api v0.19.4

replace sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.4
