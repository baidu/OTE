#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
echo "script root: $SCRIPT_ROOT"
#CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}
if [ -f go.mod ]; then
    echo "modules on"
    code_generator_path=`grep code-generator $SCRIPT_ROOT/go.mod|awk '{print $1"@"$2}'`
    CODEGEN_PKG=$GOPATH/pkg/mod/$code_generator_path
    echo "code gen path: $CODEGEN_PKG"
    chmod +x "${CODEGEN_PKG}"/generate-groups.sh
    chmod -R +w $CODEGEN_PKG
fi

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
"${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/baidu/ote-stack/pkg/generated github.com/baidu/ote-stack/pkg/apis \
  ote:v1 \
  --output-base "$SCRIPT_ROOT" \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# To use your own boilerplate text append:
#   --go-header-file "${SCRIPT_ROOT}"/hack/custom-boilerplate.go.txt

# move generated code to pkg and clean the output
cp -R $SCRIPT_ROOT/github.com/baidu/ote-stack/pkg/* $SCRIPT_ROOT/pkg && rm -rf $SCRIPT_ROOT/github.com
