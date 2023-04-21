# crdb-kubernetes
A very simple app that demonstrates the ease of scaling and upgrading a CockroachDB cluster in Kubernetes

### Setup

##### NodePort

Initialise Kubernetes:
```
k3d cluster create demo \
    -p "26257:30080@agent:0" \
    -p "8080:30081@agent:0" \
    --agents 2
```

Apply CockroachDB stateful set:
```
kubectl apply -f manifests/statefulset.yaml
kubectl apply -f manifests/init.yaml

see -n 1 kubectl get pods

kubectl apply -f manifests/ingress.yaml
```

##### LoadBalancer

Initialise Kubernetes:
```
k3d cluster create demo \
    -p "26257:80@loadbalancer"
```

Apply CockroachDB stateful set:
```
kubectl apply -f manifests/statefulset.yaml
kubectl apply -f manifests/init.yaml

see -n 1 kubectl get pods
```

### Demo basics

Connect to the Console

Connect to CockroachDB
```
cockroach sql --url "postgres://root@localhost:26257/defaultdb?sslmode=disable"
```

### Demo various operations

Kick-off a Go app
```
go run apps/ping.go
```

* Upgrade the cluster to v22.2.8
* Scale up to 5 nodes

### Kill everything
```
kubectl delete all --all -n default
kubectl delete pvc datadir-cockroachdb-0 datadir-cockroachdb-1 datadir-cockroachdb-2 datadir-cockroachdb-3 datadir-cockroachdb-4
```