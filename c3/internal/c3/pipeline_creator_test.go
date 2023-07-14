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
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestCreatePipeline(t *testing.T) {
	updatedCfg, err := addPipelineWithUnusedComponentsToExistingConfig(pipelineTypeMetrics, yamlWithUnusedReceiverAndExporter())
	require.NoError(t, err)
	fmt.Printf(": %v\n", updatedCfg)
}

func TestGetAvailablePipelineKey(t *testing.T) {
	cfg := testCfgWithUnusedReceiver(t)
	key := getAvailablePipelineKey(pipelineTypeMetrics, cfg)
	assert.Equal(t, "metrics/1", key)
}

func TestFindConfiguredComponentNames(t *testing.T) {
	m := testCfgWithUnusedReceiver(t)
	r, p, e := findConfiguredComponentNames(m)
	sort.Strings(r)
	assert.EqualValues(t, []string{"hostmetrics", "redis"}, r)
	assert.Nil(t, p)
	assert.EqualValues(t, []string{"logging"}, e)
}

func TestFindUsedComponentNames(t *testing.T) {
	m := testCfgWithUnusedReceiver(t)
	r, p, e := findUsedComponentNames(m)
	assert.EqualValues(t, []string{"hostmetrics"}, r)
	assert.Nil(t, p)
	assert.EqualValues(t, []string{"logging"}, e)
}

func TestFindUnusedComponentNames(t *testing.T) {
	m := testCfgWithUnusedReceiver(t)
	r, p, e := findUnusedComponentNames(m)
	assert.EqualValues(t, []string{"redis"}, r)
	assert.Nil(t, p)
	assert.Nil(t, e)
}

func testCfgWithUnusedReceiver(t *testing.T) map[string]any {
	yml := yamlWithUnusedReceiver()
	m := map[string]any{}
	err := yaml.Unmarshal([]byte(yml), &m)
	require.NoError(t, err)
	return m
}

func yamlWithUnusedReceiver() string {
	yml := `
receivers:
  hostmetrics:
    collection_interval: 4s
    scrapers:
      memory:
  redis:
exporters:
  logging:
service:
  pipelines:
    metrics:
      receivers:
        - hostmetrics
      exporters:
        - logging
`
	return yml
}

func yamlWithUnusedReceiverAndExporter() string {
	yml := `
receivers:
  hostmetrics:
    collection_interval: 4s
    scrapers:
      memory:
  redis:
exporters:
  logging:
  file:
    path: /foo/bar
service:
  pipelines:
    metrics:
      receivers:
        - hostmetrics
      exporters:
        - logging
`
	return yml
}
