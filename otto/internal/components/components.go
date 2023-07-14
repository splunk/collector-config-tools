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

package components

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
)

func Components() (otelcol.Factories, error) {
	var err error
	factories := otelcol.Factories{}
	factories.Extensions, err = extension.MakeFactoryMap(
		healthcheckextension.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.Receivers, err = receiver.MakeFactoryMap(
		// carbonreceiver.NewFactory(),
		// cloudfoundryreceiver.NewFactory(),
		// collectdreceiver.NewFactory(),
		// fluentforwardreceiver.NewFactory(),
		// filelogreceiver.NewFactory(),
		// hostmetricsreceiver.NewFactory(),
		// jaegerreceiver.NewFactory(),
		// journaldreceiver.NewFactory(),
		// k8sclusterreceiver.NewFactory(),
		// k8seventsreceiver.NewFactory(),
		// kafkametricsreceiver.NewFactory(),
		// kafkareceiver.NewFactory(),
		// kubeletstatsreceiver.NewFactory(),
		// mongodbatlasreceiver.NewFactory(),
		// otlpreceiver.NewFactory(),
		// prometheusexecreceiver.NewFactory(),
		// prometheusreceiver.NewFactory(),
		// receivercreator.NewFactory(),
		redisreceiver.NewFactory(),
		// sapmreceiver.NewFactory(),
		// signalfxreceiver.NewFactory(),
		// simpleprometheusreceiver.NewFactory(),
		// splunkhecreceiver.NewFactory(),
		// sqlqueryreceiver.NewFactory(),
		// statsdreceiver.NewFactory(),
		// syslogreceiver.NewFactory(),
		// tcplogreceiver.NewFactory(),
		// zipkinreceiver.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.Exporters, err = exporter.MakeFactoryMap(
		fileexporter.NewFactory(),
		// kafkaexporter.NewFactory(),
		// loggingexporter.NewFactory(),
		// otlpexporter.NewFactory(),
		// otlphttpexporter.NewFactory(),
		// sapmexporter.NewFactory(),
		// signalfxexporter.NewFactory(),
		// splunkhecexporter.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.Processors, err = processor.MakeFactoryMap(
		attributesprocessor.NewFactory(),
		// batchprocessor.NewFactory(),
		// filterprocessor.NewFactory(),
		// groupbyattrsprocessor.NewFactory(),
		// k8sattributesprocessor.NewFactory(),
		// memorylimiterprocessor.NewFactory(),
		// metricstransformprocessor.NewFactory(),
		// probabilisticsamplerprocessor.NewFactory(),
		// resourcedetectionprocessor.NewFactory(),
		// resourceprocessor.NewFactory(),
		// routingprocessor.NewFactory(),
		// spanmetricsprocessor.NewFactory(),
		// spanprocessor.NewFactory(),
		// tailsamplingprocessor.NewFactory(),
		// transformprocessor.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	return factories, nil
}
