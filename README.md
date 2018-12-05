#k2 CLI Client

This is the CLI client for managing [k2](https://github.com/vllry/k2).

##Use

Currently k2ctl is only a crude test for the gRPC API, and performs a hardcoded action.

#Building

Run `./build.sh` to ensure dep modules exist, and build the `./k2ctl` binary.

#Developing

k2ctl uses the k2 gRPC API. It requires `k2/api` as a dependency, installed with dep automatically. When making concurrent changes to the k2 API, you must point dep to a remote branch, or manually copy files (if offline).

An example Gopkg.toml change is:

```
[[constraint]]
   branch = "add-workloads"
   name = "github.com/vllry/k2"
```