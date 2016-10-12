Ergo - Errands in Go POC

First target and login (admin/admin) in to your BOSH director, then run the following commands to generate a BOSH manifest, create a dev release, deploy it and run the "deploy-service-broker" errands.

Currently only a BOSH-lite manifest is provided.

```
$ bosh target https://192.168.50.4
$ ./manifests/make-boshlite-manifest.sh
$ bosh deployment ./manifests/ergo-boshlite-manifest.yml
```

If this is your first time creating a dev release and have not added the golang archive to your local bosh blobs:

```
$ wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
$ bosh add blob go1.7.1.linux-amd64.tar.gz golang/
$ rm go1.7.1.linux-amd64.tar.gz
```

You do not need to do 'bosh upload blobs' per the output for the dev release.

```
$ bosh create release --force --name ergo
$ bosh upload release ./dev_releases/ergo/ergo-0+dev.XXXX.yml (the value for XXXX can be found in the output of the create release commend executed prior)
$ bosh -n deploy
$ bosh run errand deploy-service-broker
```

To delete the deployment:

```
$ bosh -n delete deployment ergo
```

To delete the release:

```
$ bosh -n delete release ergo
```
