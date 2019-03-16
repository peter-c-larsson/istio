// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bookinfo

import (
	"testing"

	"istio.io/istio/pkg/test/framework2/components/deployment"
	"istio.io/istio/pkg/test/framework2/core"
)

type Config struct {
	Namespace core.Namespace
}

// DeployOrFail returns a new instance of deployed BookInfo or fails test
func Deploy(ctx core.Context, cfg bookInfoConfig) (i deployment.Instance, err error) {
	err = core.UnsupportedEnvironment(ctx.Environment())

	ctx.Environment().Case(core.Kube, func() {
		i, err = deploy(ctx, cfg)
	})

	return
}

// DeployOrFail returns a new instance of deployed BookInfo or fails test
func DeployOrFail(t *testing.T, ctx core.Context, cfg bookInfoConfig) deployment.Instance {
	t.Helper()

	i, err := Deploy(ctx, cfg)
	if err != nil {
		t.Fatalf("bookinfo.DeployOrFail: %v", err)
	}

	return i
}