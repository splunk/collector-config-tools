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

package otto

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	"golang.org/x/net/websocket"
)

type metricsReceiverWrapper struct {
	receiver.Metrics
	repeater *metricsRepeater
}

func newMetricsReceiverWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory receiver.Factory) (metricsReceiverWrapper, error) {
	repeater := newMetricsRepeater(logger, ws)
	rcvr, err := factory.CreateMetricsReceiver(
		context.Background(),
		receiver.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
		repeater,
	)
	return metricsReceiverWrapper{
		Metrics:  rcvr,
		repeater: repeater,
	}, err
}

func (w metricsReceiverWrapper) setNextMetricsConsumer(next consumer.Metrics) {
	w.repeater.setNext(next)
}

func (w metricsReceiverWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type logsReceiverWrapper struct {
	receiver.Logs
	repeater *logsRepeater
}

func newLogsReceiverWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory receiver.Factory) (logsReceiverWrapper, error) {
	repeater := newLogsRepeater(logger, ws)
	rcvr, err := factory.CreateLogsReceiver(
		context.Background(),
		receiver.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
		repeater,
	)
	return logsReceiverWrapper{
		Logs:     rcvr,
		repeater: repeater,
	}, err
}

func (w logsReceiverWrapper) setNextLogsConsumer(next consumer.Logs) {
	w.repeater.setNext(next)
}

func (w logsReceiverWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type tracesReceiverWrapper struct {
	receiver.Traces
	repeater *tracesRepeater
}

func newTracesReceiverWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory receiver.Factory) (tracesReceiverWrapper, error) {
	repeater := newTracesRepeater(logger, ws)
	rcvr, err := factory.CreateTracesReceiver(
		context.Background(),
		receiver.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
		repeater,
	)
	return tracesReceiverWrapper{
		Traces:   rcvr,
		repeater: repeater,
	}, err
}

func (w tracesReceiverWrapper) setNextTracesConsumer(next consumer.Traces) {
	w.repeater.setNext(next)
}

func (w tracesReceiverWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type metricsProcessorWrapper struct {
	processor.Metrics
	repeater *metricsRepeater
}

func newMetricsProcessorWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory processor.Factory) (metricsProcessorWrapper, error) {
	repeater := newMetricsRepeater(logger, ws)
	proc, err := factory.CreateMetricsProcessor(
		context.Background(),
		processor.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
		repeater,
	)
	return metricsProcessorWrapper{
		Metrics:  proc,
		repeater: repeater,
	}, err
}

func (w metricsProcessorWrapper) setNextMetricsConsumer(next consumer.Metrics) {
	w.repeater.setNext(next)
}

func (w metricsProcessorWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type logsProcessorWrapper struct {
	processor.Logs
	repeater *logsRepeater
}

func newLogsProcessorWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory processor.Factory) (logsProcessorWrapper, error) {
	repeater := newLogsRepeater(logger, ws)
	proc, err := factory.CreateLogsProcessor(
		context.Background(),
		processor.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
		repeater,
	)
	return logsProcessorWrapper{
		Logs:     proc,
		repeater: repeater,
	}, err
}

func (w logsProcessorWrapper) setNextLogsConsumer(next consumer.Logs) {
	w.repeater.setNext(next)
}

func (w logsProcessorWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type tracesProcessorWrapper struct {
	processor.Traces
	repeater *tracesRepeater
}

func newTracesProcessorWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory processor.Factory) (tracesProcessorWrapper, error) {
	repeater := newTracesRepeater(logger, ws)
	proc, err := factory.CreateTracesProcessor(
		context.Background(),
		processor.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
		repeater,
	)
	return tracesProcessorWrapper{
		Traces:   proc,
		repeater: repeater,
	}, err
}

func (w tracesProcessorWrapper) setNextTracesConsumer(next consumer.Traces) {
	w.repeater.setNext(next)
}

func (w tracesProcessorWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type metricsExporterWrapper struct {
	exporter.Metrics
	repeater *metricsRepeater
}

func newMetricsExporterWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory exporter.Factory) (metricsExporterWrapper, error) {
	repeater := newMetricsRepeater(logger, ws)
	ex, err := factory.CreateMetricsExporter(
		context.Background(),
		exporter.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
	)
	return metricsExporterWrapper{
		Metrics:  ex,
		repeater: repeater,
	}, err
}

func (w metricsExporterWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type logsExporterWrapper struct {
	exporter.Logs
	repeater *logsRepeater
}

func newLogsExporterWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory exporter.Factory) (logsExporterWrapper, error) {
	ex, err := factory.CreateLogsExporter(
		context.Background(),
		exporter.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
	)
	return logsExporterWrapper{
		Logs:     ex,
		repeater: newLogsRepeater(logger, ws),
	}, err
}

func (w logsExporterWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}

type tracesExporterWrapper struct {
	exporter.Traces
	repeater *tracesRepeater
}

func newTracesExporterWrapper(logger *log.Logger, ws *websocket.Conn, cfg component.Config, factory exporter.Factory) (tracesExporterWrapper, error) {
	ex, err := factory.CreateTracesExporter(
		context.Background(),
		exporter.CreateSettings{
			TelemetrySettings: componenttest.NewNopTelemetrySettings(),
		},
		cfg,
	)
	return tracesExporterWrapper{
		Traces:   ex,
		repeater: newTracesRepeater(logger, ws),
	}, err
}

func (w tracesExporterWrapper) waitForStopMessage() {
	w.repeater.waitForStopMessage()
}
