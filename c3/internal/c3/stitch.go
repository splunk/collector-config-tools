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

type stitchHandler struct {
	logger *log.Logger
}

func (h stitchHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		h.logger.Printf("stitchHandler: ServeHTTP: error reading request: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	m := map[string]any{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		h.logger.Printf("stitchHandler: ServeHTTP: error unmarshaling JSON: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	stitched, err := stitch(
		m["collectorYaml"].(string),
		m["componentGroup"].(string),
		m["componentName"].(string),
		m["componentCfg"],
	)
	if err != nil {
		h.logger.Printf("stitchHandler: ServeHTTP: error unmarshaling JSON: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	stitchedYaml, err := toYaml(stitched)
	if err != nil {
		h.logger.Printf("stitchHandler: ServeHTTP: error unmarshaling JSON: %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(stitchedYaml)
}

type ccfg struct {
	Receivers  any `yaml:"receivers,omitempty"`
	Processors any `yaml:"processors,omitempty"`
	Exporters  any `yaml:"exporters,omitempty"`
	Service    any `yaml:"service,omitempty"`
}

func stitch(yml, componentGroup, componentType string, newComponentCfg any) (ccfg, error) {
	stitchedConfig := map[string]any{}
	err := yaml.Unmarshal([]byte(yml), &stitchedConfig)
	if err != nil {
		return ccfg{}, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}
	pluralGroup := componentGroup + "s" // e.g. "receivers"
	var groupComponents map[string]any
	groupComponents, ok := stitchedConfig[pluralGroup].(map[string]any)
	if !ok {
		groupComponents = map[string]any{}
		stitchedConfig[pluralGroup] = groupComponents
	}
	groupComponents[componentType] = newComponentCfg
	out := ccfg{
		Receivers:  stitchedConfig["receivers"],
		Processors: stitchedConfig["processors"],
		Exporters:  stitchedConfig["exporters"],
		Service:    stitchedConfig["service"],
	}
	return out, nil
}
