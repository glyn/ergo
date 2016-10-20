# Ergo - Errands in Go POC

## bosh release development

First target and login (admin/admin) in to your BOSH director, then run the following commands to generate a BOSH manifest, create a dev release, deploy it and run the "deploy-service-broker" errands.

Currently only a BOSH-lite manifest is provided.

```
$ bosh target https://192.168.50.4
$ ./manifests/make-boshlite-manifest.sh
$ bosh deployment ./manifests/ergo-boshlite-manifest.yml
```

If this is your first time creating a dev release, add the golang archive and the cf CLI to your local bosh blobs, for example:

```
$ wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
$ bosh add blob go1.7.1.linux-amd64.tar.gz golang/
$ rm go1.7.1.linux-amd64.tar.gz
$ wget https://s3-us-west-1.amazonaws.com/cf-cli-releases/releases/v6.22.1/cf-cli_6.22.1_linux_x86-64.tgz
$ bosh add blob cf-cli_6.22.1_linux_x86-64.tgz cf-cli/
$ rm cf-cli_6.22.1_linux_x86-64.tgz
```

You may also need to upload a stemcell to BOSH-lite, for example:

```
$ bosh public stemcells
$ bosh download public stemcell bosh-stemcell-389-warden-boshlite-ubuntu-trusty-go_agent.tgz
$ bosh upload stemcell bosh-stemcell-389-warden-boshlite-ubuntu-trusty-go_agent.tgz
```

You do not need to do 'bosh upload blobs' per the output for the dev release.

```
$ bosh create release --force --name ergo
$ bosh upload release ./dev_releases/ergo/ergo-0+dev.XXXX.yml (the value for XXXX can be found in the output of the create release commend executed prior)
$ bosh -n deploy
$ bosh run errand deploy-service-broker
```

Observe the following evidence of the errand executing Go code in the output:

```
[stdout]
* Starting deploy errand
* Finished deploy errand
```

To delete the deployment:

```
$ bosh -n delete deployment ergo
```

To delete the release:

```
$ bosh -n delete release ergo
```

## Go errand development

* Get this repository into your $GOPATH:

```
go get github.com/glyn/ergo
```
Note: this will print a warning `can't load package: package github.com/glyn/ergo: no buildable Go source files [...]` which you can ignore. This is because there are no Go files in the root of the repository.

Then change directory to the repository, typically to $GOPATH/src/github.com/glyn/ergo.

Adjust the remote if you need to be able to push back changes and check out the branch you want to work on:
```
git remote set-url origin git@github.com:glyn/ergo.git
git checkout -b feature-branch-nnnnnnnn
```

* Make code changes as needed then run the unit tests:
```
go test ./...
```

Observe that the tests pass, for example:
```
$ go test ./...
ok  	github.com/glyn/ergo/cf	0.072s
?   	github.com/glyn/ergo/deploy-errand	[no test files]
```

* If you want to test an errand binary locally, build and install, for example:

```
go install github.com/glyn/ergo/deploy-errand
```

The resulting binary can be found in the usual place, typically `$GOPATH/bin`.

This binary will be compiled for the architecture you are running on and you can run it as usual:

```
./deploy-errand
```
