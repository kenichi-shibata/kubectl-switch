# kube-env
kubectl version switcher

Why?
-------

Because Kubectl Version Skew https://kubernetes.io/docs/setup/version-skew-policy/


What?
-----
Get an asciicinema here 

Install 
--------
go get -u github.com/kenichi-shibata/kube-env

Usage
-------
```
go run kubectl-switch.go
go run kubectl-switch.go #automatically changes version
# list available versions
ls ~/.kube/kubectl/

Config 
-------

This creates a config file at `~/.kube/kubectl/config` if its not created already. Otherwise it will read these values

```
{
 "url_prefix": "https://storage.googleapis.com/kubernetes-release/release",
 "version": "v1.14.3"
}
```

Setup draft for local dev with a kube cluster
-----------

Install draft 
```
brew install azure/draft/draft 
```

Initialize draft
```
draft init # Initializes in ~/.draft
draft create # created the draft dockerfile, chart and other artifacts to deploy to kubernetes
draft config set registry docker.io/kenichishibata # set the docker registry
draft up
```
