#!/bin/bash -e
# Copyright 2020 Celo Org
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


BRANCH=${1:-master}
SCRIPTS_DIR=`dirname $0`
cd $SCRIPTS_DIR/../..

LOCATION="./external/bls-zexe"
if [[ -e $LOCATION ]]; then
  git rm -r $LOCATION
  git commit -m"bls-zexe: removes old version"
fi
git subtree add --prefix external/bls-zexe https://github.com/celo-org/bls-zexe $BRANCH --squash
