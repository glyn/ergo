#!/bin/sh

set -e

MANIFEST_NAME="ergo-boshlite-manifest.yml"
SCRIPT_HOME=$(cd `dirname "${BASH_SOURCE[0]}"` && pwd)

cat $SCRIPT_HOME/$MANIFEST_NAME.tpl | sed -e "s/%%UUID%%/$(bosh status --uuid)/g" > $SCRIPT_HOME/$MANIFEST_NAME

echo "Generated manifest: $SCRIPT_HOME/ergo-boshlite-manifest.yml"
