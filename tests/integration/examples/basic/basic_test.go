//  Copyright 2019 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package basic contains an example test suite for showcase purposes.
package basic

import (
	"testing"

	"istio.io/istio/pkg/test/framework2"
	"istio.io/istio/pkg/test/framework2/components/galley"
	"istio.io/istio/pkg/test/framework2/components/mixer"
	"istio.io/istio/pkg/test/framework2/runtime"
)

// To opt-in to the test framework, implement a TestMain, and call test.RunSuite.
func TestMain(m *testing.M) {
	framework2.RunSuite("basic_test", m, nil)
}

func TestBasic(t *testing.T) {
	framework2.Run(t, func(s *runtime.TestContext) {

		gl := galley.NewOrFail(t, s)
		mx := mixer.NewOrFail(s, gl)

		gl.ApplyConfig(`
apiVersion: "config.istio.io/v1alpha2"
kind: metric
metadata:
 name: metric1
spec:
 value: "2"
 dimensions:
   source: source.name | "mysrc"
   target_ip: destination.name | "mytarget"`)

		// As an example, the following method calls the Report operation against Mixer's own API directly.
		mx.Report(t, map[string]interface{}{})
	})
}

//func TestBasic(t *testing.T) {
//	// Call Requires to explicitly initialize dependencies that the test needs.
//	ctx := framework.GetContext(t)
//	ctx.RequireOrSkip(t, lifecycle.Test, &ids.Mixer)
//
//	// Call environment.Configure to set Istio-wide configuration for the test.
//	mixer := components.GetMixer(ctx, t)
//	mixer.Configure(t, lifecycle.Test, `
//apiVersion: "config.istio.io/v1alpha2"
//kind: metric
//metadata:
// name: metric1
//spec:
// value: "2"
// dimensions:
//   source: source.name | "mysrc"
//   target_ip: destination.name | "mytarget"
//`)
//
//	// As an example, the following method calls the Report operation against Mixer's own API directly.
//	mixer.Report(t, map[string]interface{}{})
//}