# kubectl-switch
kubectl version switcher

Why?
-------

Because Kubectl Version Skew https://kubernetes.io/docs/setup/version-skew-policy/

If you work with two clusters with at least 2 minor versions apart then there will be functionalities that will be missing. You need to switch your version.

Kubectl-Switch to the rescue

What?
-----
Get an asciicinema here 

Install 
--------
go get -u github.com/kenichi-shibata/kubectl-switch

Usage
-------
```
go run kubectl-switch.go
go run kubectl-switch.go #automatically changes version
# list available versions
ls ~/.kube/kubectl/
```
Config 
-------

This creates a config file at `~/.kube/kubectl/config` if its not created already. Otherwise it will read these values

```
{
 "url_prefix": "https://storage.googleapis.com/kubernetes-release/release",
 "version": "v1.14.3"
}
```

