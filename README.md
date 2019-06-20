# kubectl-switch
kubectl version switcher

Why?
-------

Because Kubectl Version Skew https://kubernetes.io/docs/setup/version-skew-policy/

If you work with two clusters with at least 2 minor versions apart then there will be functionalities that will be missing. You need to switch your version.

Kubectl-Switch to the rescue

What?
-----


Install 
--------
go get -u github.com/kenichi-shibata/kubectl-switch

Usage
-------
```
# downloads v1.14.3 or -k <version> or from config file ~/.kube/kubectl/config
kubectl-switch download 
kubectl-switch -k v.1.11.9
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

Alternatives
------------
* [asdf](https://asdf-vm.com/#/) with [kubectl-plugin](https://github.com/Banno/asdf-kubectl)