#!/bin/bash

# Copyright 2019 The Kubernetes Authors.
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

set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
cd "${REPO_ROOT}" || exit 1

# shellcheck source=../hack/ensure-go.sh
source "${REPO_ROOT}/hack/ensure-go.sh"
# shellcheck source=../hack/ensure-kind.sh
source "${REPO_ROOT}/hack/ensure-kind.sh"
# shellcheck source=../hack/ensure-kubectl.sh
source "${REPO_ROOT}/hack/ensure-kubectl.sh"
# shellcheck source=../hack/ensure-kustomize.sh
source "${REPO_ROOT}/hack/ensure-kustomize.sh"

make test-e2e
test_status="${?}"

# TODO last chance to clean up resources if prow job leaves something behind

exit "${test_status}"

export CLIENT_ID=da8c7265-60fc-43bc-91b5-0aa0f1b05f25
export CLIENT_SECRET=cuVJ7.i60I??PhwwbqI-scrfXfg6PHd:
export SUBSCRIPTION_ID=e3cdcc6a-93e3-4bf1-b2d3-54dba53a5764
export TENANT_ID=javierdarsieoutlook.onmicrosoft.com