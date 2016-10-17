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
* Starting deploy errand
....
....
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

Setting up your local environmet to develop errand code and run tests:

* First append to your $GOPATH the directory housing the Go errands:

```
export GOPATH=$GOPATH:/path/to/ergo/errands/
```

* NOTE: Consider modifying your shell environment configuration (.bash_profile, etc) file so you don't need to modify your GOPATH each time.

* Make code changes as needed then build and install, for example:

```
go install deploy-errand
```

The resulting binary can be found in:

```
errands/bin/deploy-errand
```

This binary will be compiled for the architecture you are running on. To invoke the errand from the command line, a few environment variables are required:

```
SYSTEM_DOMAIN - the system domain that will be used to target CF, i.e.: lime.springapps.io
ADMIN_USER - the CF admin user
ADMIN_PASSWORD - the CF admin password
```

The following environment variables are optional can be set to customize settings for local development:

```
VCAP_USER_NAME - the username to override the default 'vcap' with
VCAP_GROUP_NAME - the group name to override the default 'vcap' with
VCAP_DIR_PREFIX - the prefix to use when creating the vcap directory structure
```

An example local invocation of the deploy-errand binary would look like:

```
SYSTEM_DOMAIN=bosh-lite.com ADMIN_USER=admin ADMIN_PASSWORD=admin VCAP_DIR_PREFIX=/tmp VCAP_USER_NAME=cschaefer VCAP_GROUP_NAME=wheel ./deploy-errand
```

You can also set these environment variables into your shell for one time usage or persist them in the appropriate shell config file

* To run tests:

```
$ cd errands/src/deploy-errand
$ go test -v
```

Observe the following evidence of the Go test in the output:

```
=== RUN   TestValidGetUserId
--- PASS: TestValidGetUserId (0.00s)
=== RUN   TestInvalidGetUserId
Failed to obtain UID for user name: user: unknown user someuserthatdoesnotexist
--- FAIL: TestInvalidGetUserId (0.00s)
	deploy_errand_test.go:25: Failed to obtain UID: user: unknown user someuserthatdoesnotexist
FAIL
exit status 1
FAIL	deploy-errand	0.007s
```
