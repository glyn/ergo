#!/bin/bash
set -e -x

export DEPLOY_ERRAND_PACKAGE=$(readlink -nf /var/vcap/packages/deploy-errand)
export CF_CLI_PACKAGE=$(readlink -nf /var/vcap/packages/cf-cli)
export PATH=$DEPLOY_ERRAND_PACKAGE/bin:$CF_CLI_PACKAGE/bin:$PATH

SCRIPT_DIR=$(dirname $0)

exec deploy-errand $SCRIPT_DIR/props.json
