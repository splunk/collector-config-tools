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
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type pipelineType string

const (
	pipelineTypeMetrics pipelineType = "metrics"
	pipelineTypeLogs    pipelineType = "logs"
	pipelineTypeTraces  pipelineType = "traces"
)

type createPipelineHandler struct {
	logger *log.Logger
}

func (h createPipelineHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		h.logger.Printf("createPipelineHandler: ServeHTTP: error reading request: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	m := map[string]string{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		h.logger.Printf("createPipelineHandler: ServeHTTP: error unmarshaling JSON: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	cfg, err := addPipelineWithUnusedComponentsToExistingConfig(pipelineType(m["pipelineType"]), m["collectorYaml"])
	if err != nil {
		h.logger.Printf("createPipelineHandler: ServeHTTP: error adding pipeline: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	yml, err := toYaml(cfg)
	if err != nil {
		h.logger.Printf("createPipelineHandler: ServeHTTP: error converting JSON to YAML: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = resp.Write(yml)
	if err != nil {
		h.logger.Printf("createPipelineHandler: ServeHTTP: error writing response: %v", err)
	}
}

func addPipelineWithUnusedComponentsToExistingConfig(pipelineType pipelineType, cfgYaml string) (ccfg, error) {
	cfg := map[string]any{}
	err := yaml.Unmarshal([]byte(cfgYaml), &cfg)
	if err != nil {
		return ccfg{}, err
	}
	r, p, e := findUnusedComponentNames(cfg)
	rpe := createPipelineRPEMap(r, p, e)
	pipelineKey := getAvailablePipelineKey(pipelineType, cfg)
	addPipeline(cfg, pipelineKey, rpe)
	out := ccfg{
		Receivers:  cfg["receivers"],
		Processors: cfg["processors"],
		Exporters:  cfg["exporters"],
		Service:    cfg["service"],
	}
	return out, nil
}

func addPipeline(cfg map[string]any, pipelineKey string, rpe map[string]any) {
	service, ok := cfg["service"].(map[string]any)
	if !ok {
		service = map[string]any{}
		cfg["service"] = service
	}
	pipelines, ok := service["pipelines"].(map[string]any)
	if !ok {
		pipelines = map[string]any{}
		service["pipelines"] = pipelines
	}
	pipelines[pipelineKey] = rpe
}

func createPipelineRPEMap(r, p, e []string) map[string]any {
	out := map[string]any{}
	if r != nil {
		out["receivers"] = r
	}
	if p != nil {
		out["processors"] = p
	}
	if e != nil {
		out["exporters"] = e
	}
	return out
}

func getAvailablePipelineKey(pipelineType pipelineType, m map[string]any) string {
	ptStr := string(pipelineType)
	service, ok := m["service"].(map[string]any)
	if !ok {
		return ptStr
	}
	pipelines, ok := service["pipelines"].(map[string]any)
	if !ok {
		return ptStr
	}
	_, ok = pipelines[ptStr]
	if !ok {
		return ptStr
	}
	i := 1 // FIXME
	return fmt.Sprintf("%s/%d", pipelineType, i)
}

func findUnusedComponentNames(cfg map[string]any) ([]string, []string, []string) {
	rc, pc, ec := findConfiguredComponentNames(cfg)
	ru, pu, eu := findUsedComponentNames(cfg)
	return subtractStringSlices(rc, ru), subtractStringSlices(pc, pu), subtractStringSlices(ec, eu)
}

func findConfiguredComponentNames(cfg map[string]any) ([]string, []string, []string) {
	rm, _ := cfg["receivers"].(map[string]any)
	pm, _ := cfg["processors"].(map[string]any)
	em, _ := cfg["exporters"].(map[string]any)
	return mapKeys(rm), mapKeys(pm), mapKeys(em)
}

func findUsedComponentNames(cfg map[string]any) ([]string, []string, []string) {
	service, ok := cfg["service"].(map[string]any)
	if !ok {
		return nil, nil, nil
	}
	pipelines, ok := service["pipelines"].(map[string]any)
	if !ok {
		return nil, nil, nil
	}
	metricsPipeline, ok := pipelines["metrics"].(map[string]any)
	if !ok {
		return nil, nil, nil
	}
	r, _ := metricsPipeline["receivers"].([]any)
	p, _ := metricsPipeline["processors"].([]any)
	e, _ := metricsPipeline["exporters"].([]any)
	return anySlicetoStringSlice(r), anySlicetoStringSlice(p), anySlicetoStringSlice(e)
}

func anySlicetoStringSlice(a []any) []string {
	var out []string
	for _, v := range a {
		out = append(out, v.(string))
	}
	return out
}

func mapKeys(m map[string]any) []string {
	var out []string
	for k := range m {
		out = append(out, k)
	}
	return out
}

func subtractStringSlices(big, little []string) []string {
	lookup := map[string]bool{}
	for _, str := range little {
		lookup[str] = true
	}
	var out []string
	for _, str := range big {
		if !lookup[str] {
			out = append(out, str)
		}
	}
	return out
}
