#!/bin/bash -e

BRANCH=${1:-master}
SCRIPTS_DIR=`dirname $0`
cd $SCRIPTS_DIR/../..

LOCATION="./external/bls-zexe"
if [[ -e $LOCATION ]]; then
  git rm -r $LOCATION
  git commit -m"bls-zexe: removes old version"
fi
git subtree add --prefix external/bls-zexe https://github.com/celo-org/bls-zexe $BRANCH --squash
