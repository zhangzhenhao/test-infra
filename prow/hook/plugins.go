/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hook

// All enabled plugins. We need to empty import them like this so that they
// will be linked into any hook binary.
import (
	_ "k8s.io/test-infra/prow/plugins/assign"
	_ "k8s.io/test-infra/prow/plugins/cla"
	_ "k8s.io/test-infra/prow/plugins/close"
	_ "k8s.io/test-infra/prow/plugins/golint"
	_ "k8s.io/test-infra/prow/plugins/heart"
	_ "k8s.io/test-infra/prow/plugins/hold"
	_ "k8s.io/test-infra/prow/plugins/label"
	_ "k8s.io/test-infra/prow/plugins/lgtm"
	_ "k8s.io/test-infra/prow/plugins/releasenote"
	_ "k8s.io/test-infra/prow/plugins/reopen"
	_ "k8s.io/test-infra/prow/plugins/shrug"
	_ "k8s.io/test-infra/prow/plugins/size"
	_ "k8s.io/test-infra/prow/plugins/slackevents"
	_ "k8s.io/test-infra/prow/plugins/trigger"
	_ "k8s.io/test-infra/prow/plugins/updateconfig"
	_ "k8s.io/test-infra/prow/plugins/wip"
	_ "k8s.io/test-infra/prow/plugins/yuks"
)
