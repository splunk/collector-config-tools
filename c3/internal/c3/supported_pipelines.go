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
	"reflect"
	"sort"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
)

type component struct {
	Name                  string
	Metrics, Traces, Logs bool
}

type components struct {
	Receivers, Processors, Exporters []component
}

func (c components) found(componentType, name string) bool {
	var typeComponents []component
	switch componentType {
	case "receiver":
		typeComponents = c.Receivers
	case "processor":
		typeComponents = c.Processors
	case "exporter":
		typeComponents = c.Exporters
	}
	for _, c := range typeComponents {
		if c.Name == name {
			return true
		}
	}
	return false
}

func factoriesToComponentTypes(factories otelcol.Factories) components {
	out := components{}
	for name, factory := range factories.Receivers {
		sp := receiverSupportedPipelines(factory)
		c := component{
			Name:    string(name),
			Metrics: sp.metrics,
			Traces:  sp.traces,
			Logs:    sp.logs,
		}
		out.Receivers = append(out.Receivers, c)
	}
	sort.Slice(out.Receivers, func(i, j int) bool {
		return out.Receivers[i].Name < out.Receivers[j].Name
	})
	for name, factory := range factories.Processors {
		sp := processorSupportedPipelines(factory)
		c := component{
			Name:    string(name),
			Metrics: sp.metrics,
			Traces:  sp.traces,
			Logs:    sp.logs,
		}
		out.Processors = append(out.Processors, c)
	}
	sort.Slice(out.Processors, func(i, j int) bool {
		return out.Processors[i].Name < out.Processors[j].Name
	})
	for name, factory := range factories.Exporters {
		sp := exporterSupportedPipelines(factory)
		c := component{
			Name:    string(name),
			Metrics: sp.metrics,
			Traces:  sp.traces,
			Logs:    sp.logs,
		}
		out.Exporters = append(out.Exporters, c)
	}
	sort.Slice(out.Exporters, func(i, j int) bool {
		return out.Exporters[i].Name < out.Exporters[j].Name
	})
	return out
}

type supportedPipelines struct {
	metrics bool
	traces  bool
	logs    bool
}

func receiverSupportedPipelines(fact receiver.Factory) supportedPipelines {
	out := supportedPipelines{}
	factV := reflect.ValueOf(fact).Elem()
	if !factV.FieldByName("CreateMetricsFunc").IsNil() {
		out.metrics = true
	}
	if !factV.FieldByName("CreateTracesFunc").IsNil() {
		out.traces = true
	}
	if !factV.FieldByName("CreateLogsFunc").IsNil() {
		out.logs = true
	}
	return out
}

func processorSupportedPipelines(fact processor.Factory) supportedPipelines {
	out := supportedPipelines{}
	factV := reflect.ValueOf(fact).Elem()
	if !factV.FieldByName("CreateMetricsFunc").IsNil() {
		out.metrics = true
	}
	if !factV.FieldByName("CreateTracesFunc").IsNil() {
		out.traces = true
	}
	if !factV.FieldByName("CreateLogsFunc").IsNil() {
		out.logs = true
	}
	return out
}

func exporterSupportedPipelines(fact exporter.Factory) supportedPipelines {
	out := supportedPipelines{}
	factV := reflect.ValueOf(fact).Elem()
	if !factV.IsValid() {
		return out
	}
	if f := factV.FieldByName("CreateMetricsFunc"); f.IsValid() && !f.IsNil() {
		out.metrics = true
	}
	if f := factV.FieldByName("CreateTracesFunc"); f.IsValid() && !f.IsNil() {
		out.traces = true
	}
	if f := factV.FieldByName("CreateLogsFunc"); f.IsValid() && !f.IsNil() {
		out.logs = true
	}
	return out
}
