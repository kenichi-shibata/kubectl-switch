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
go run kubectl-switch.go v1.11.2
go run kubectl-switch.go #automatically changes version
# list available versions
ls ~/.kube/kubectl/
```
