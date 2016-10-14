Ergo - Errands in Go POC

First target and login (admin/admin) in to your BOSH director, then run the following commands to generate a BOSH manifest, create a dev release, deploy it and run the "deploy-service-broker" errands.

Currently only a BOSH-lite manifest is provided. The provided manifest also assumes CF is deployed to BOSH-lite. If this is not the case, edit:

```
ergo-boshlite-manifest.yml.tpl
```

and modify:

```
properties.domain
properties.spring_cloud_broker.cf.admin_user
properties.spring_cloud_broker.cf.admin_password
```

To whatever values that match the CF environment you are targeting. Information on deploying CF to bosh-lite can be found at: https://github.com/cloudfoundry/bosh-lite#deploy-cloud-foundry

```
$ bosh target https://192.168.50.4
$ ./manifests/make-boshlite-manifest.sh
$ bosh deployment ./manifests/ergo-boshlite-manifest.yml
```

If this is your first time creating a dev release and have not added the golang archive to your local bosh blobs, for example:

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
hello from golang
```

To delete the deployment:

```
$ bosh -n delete deployment ergo
```

To delete the release:

```
$ bosh -n delete release ergo
```
