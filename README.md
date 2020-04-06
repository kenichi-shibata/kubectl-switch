# kubectl-switch
kubectl version switcher

Quickstart
--------

```
$ go get -v -u github.com/kenichi-shibata/kubectl-switch/cmd
$ kubectl-switch version
v0.0.2
$ kubectl-switch download -k v1.15.11
$ kubectl version --client # to verify

```



Why?
-------

Because Kubectl Version Skew https://kubernetes.io/docs/setup/version-skew-policy/

If you work with two clusters with at least 2 minor versions apart then there will be functionalities that will be missing. You need to switch your version.

Kubectl-Switch to the rescue

What?
-----
[![asciicast](https://asciinema.org/a/rNUZ5ywLkNdAXnj3GtQBlIvtf.svg)](https://asciinema.org/a/rNUZ5ywLkNdAXnj3GtQBlIvtf)

Install
--------

```
git clone git@github.com:kenichi-shibata/kubectl-switch
go build .
./kubectl-switch
```

Usage
-------
```
kubectl-switch download # get the latest stable version
kubectl-switch download -k v.1.11.9 # switch to verison v1.11.9
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

The binares will created in `~/.kube/kubectl/` and the main symlink will be in `~/.kube/kubectl/kubectl` which will be symlinked to the active version.

Alternatives
------------
* [asdf](https://asdf-vm.com/#/) with [kubectl-plugin](https://github.com/Banno/asdf-kubectl)

Generate Supported Versions
--------------
```
curl -s https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=100 | jq .[].Name > supported_versions
```

List of all supported versions

https://raw.githubusercontent.com/kenichi-shibata/kubectl-switch/master/supported_versions
