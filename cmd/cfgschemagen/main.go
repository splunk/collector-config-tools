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

package main

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/configschema"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/zpagesextension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/multierr"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

const module = "github.com/splunk/collector-config-tools/cmd/cfgschemagen"

func run() error {
	factories, err := components()
	if err != nil {
		return err
	}
	return configschema.GenerateYAMLFiles(factories, ".", "outputDir", module)
}

func components() (otelcol.Factories, error) {
	var errs []error
	extensions, err := extension.MakeFactoryMap(
		zpagesextension.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	receivers, err := receiver.MakeFactoryMap(redisreceiver.NewFactory())
	if err != nil {
		errs = append(errs, err)
	}

	processors, err := processor.MakeFactoryMap(filterprocessor.NewFactory())
	if err != nil {
		errs = append(errs, err)
	}

	exporters, err := exporter.MakeFactoryMap(otlpexporter.NewFactory())
	if err != nil {
		errs = append(errs, err)
	}

	return otelcol.Factories{
		Extensions: extensions,
		Receivers:  receivers,
		Processors: processors,
		Exporters:  exporters,
	}, multierr.Combine(errs...)
}
