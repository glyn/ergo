Ergo - Errands in Go POC

First target and login in to your BOSH director, then run the following commands to generate a BOSH manifest, create a dev release, deploy it and run the "deploy-service-broker" errands.

Currently only a BOSH-lite manifest is provided.

```
$ ./manifests/make-boshlite-manifest.sh
$ bosh deployment ./manifests/ergo-boshlite-manifest.yml
$ bosh create release --force --name ergo
$ bosh upload release ./dev_releases/ergo/ergo-0+dev.XXXX.yml (the value for XXXX can be found in the output of the create release commend executed prior)
$ bosh -n deploy
$ bosh run errand deploy-service-broker
```

