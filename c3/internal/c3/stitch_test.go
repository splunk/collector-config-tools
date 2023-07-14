// Copyright Splunk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package c3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestStitch(t *testing.T) {
	originalCfg := map[string]any{
		"exporters": map[string]any{
			"file": map[string]any{
				"path": "/foo",
			},
		},
	}
	originalYaml, err := yaml.Marshal(originalCfg)
	require.NoError(t, err)
	stitchedCfg, err := stitch(
		string(originalYaml),
		"receiver",
		"redis",
		map[string]any{
			"collection_interval": "42s",
		},
	)
	require.NoError(t, err)
	fmt.Printf("cfg: %v\n", stitchedCfg)
	stitchedYaml, err := yaml.Marshal(stitchedCfg)
	require.NoError(t, err)
	println(string(stitchedYaml))
}
