# crdb-kubernetes
A very simple app that demonstrates the ease of scaling and upgrading a CockroachDB cluster in Kubernetes

### Setup

Initialise Kubernetes (I'm using [kind](https://kind.sigs.k8s.io)):
```
$ kind create cluster
```

Apply CockroachDB stateful set
```
$ kubectl apply -f manifests/statefulset.yaml
$ kubectl get pods

$ kubectl apply -f manifests/init.yaml
$ kubectl get pods
```

```
$ kubectl port-forward svc/cockroachdb-public 26257:26257 8080:8080
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