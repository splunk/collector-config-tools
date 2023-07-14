// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configschema // import "github.com/open-telemetry/opentelemetry-collector-contrib/lib/configschema"

import (
	"fmt"
	"reflect"

	"go.opentelemetry.io/collector/otelcol"
	"gopkg.in/yaml.v2"
)

// generateYAMLFiles is the entry point for cfgschema. Component factories are
// passed in so it can be used by other distros.
func generateYAMLFiles(factories otelcol.Factories, commentReader commentReaderFunc, writeYaml writeYamlFunc) error {
	configs := getAllCfgInfos(factories)
	for _, ci := range configs {
		err := writeComponentYAML(writeYaml, ci, commentReader)
		if err != nil {
			fmt.Printf("skipped writing config meta yaml for type: %s: %v\n", ci.Type, err)
		}
	}
	return nil
}

func writeComponentYAML(writeYaml writeYamlFunc, ci cfgInfo, commentReader commentReaderFunc) error {
	fields, err := readFields(reflect.ValueOf(ci.CfgInstance), commentReader)
	if err != nil {
		return fmt.Errorf("error reading fields for component for type: %s: %w", ci.Type, err)
	}
	yamlBytes, err := yaml.Marshal(fields)
	if err != nil {
		return fmt.Errorf("error marshaling to yaml for type: %s: %w", ci.Type, err)
	}
	err = writeYaml(ci, yamlBytes)
	if err != nil {
		return fmt.Errorf("error writing component yaml for type: %s: %w", ci.Type, err)
	}
	return nil
}
