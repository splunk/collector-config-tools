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
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/receiver"
	"golang.org/x/net/websocket"
)

type receiverSocketHandler struct {
	logger          *log.Logger
	pipeline        *pipeline
	receiverFactory receiver.Factory
}

func (h receiverSocketHandler) handle(ws *websocket.Conn) {
	err := h.doHandle(ws)
	if err != nil {
		sendErr(ws, h.logger, "receiverSocketHandler", err)
	}
}

func (h receiverSocketHandler) doHandle(ws *websocket.Conn) error {
	pipelineType, conf, err := readSocket(ws)
	if err != nil {
		return err
	}
	receiverConfig := h.receiverFactory.CreateDefaultConfig()
	conf.Unmarshal(receiverConfig, confmap.WithIgnoreUnused())
	if err != nil {
		return err
	}
	switch pipelineType {
	case "metrics":
		h.startMetricsReceiver(ws, receiverConfig)
	case "logs":
		h.startLogsReceiver(ws, receiverConfig)
	case "traces":
		h.startTracesReceiver(ws, receiverConfig)
	}
	return nil
}

func (h receiverSocketHandler) startMetricsReceiver(
	ws *websocket.Conn,
	cfg component.Config,
) {
	wrapper, err := newMetricsReceiverWrapper(h.logger, ws, cfg, h.receiverFactory)
	if err != nil {
		sendErr(ws, h.logger, "failed to create metrics receiver wrapper", err)
		return
	}
	h.pipeline.connectMetricsReceiverWrapper(wrapper)
	err = wrapper.Start(context.Background(), componenttest.NewNopHost())
	if err != nil {
		sendErr(ws, h.logger, "failed to start metrics receiver", err)
		return
	}
	wrapper.waitForStopMessage()
	err = wrapper.Shutdown(context.Background())
	if err != nil {
		sendErr(ws, h.logger, "failed to shut down metrics receiver", err)
		return
	}
	h.pipeline.disconnectMetricsReceiverWrapper()
}

func (h receiverSocketHandler) startLogsReceiver(
	ws *websocket.Conn,
	cfg component.Config,
) {
	wrapper, err := newLogsReceiverWrapper(h.logger, ws, cfg, h.receiverFactory)
	if err != nil {
		sendErr(ws, h.logger, "failed to create logs receiver wrapper", err)
		return
	}
	h.pipeline.connectLogsReceiverWrapper(wrapper)
	err = wrapper.Start(context.Background(), componenttest.NewNopHost())
	if err != nil {
		sendErr(ws, h.logger, "failed to start logs receiver", err)
		return
	}
	wrapper.waitForStopMessage()
	err = wrapper.Shutdown(context.Background())
	if err != nil {
		sendErr(ws, h.logger, "failed to shut down logs receiver", err)
		return
	}
	h.pipeline.disconnectLogsReceiverWrapper()
}

func (h receiverSocketHandler) startTracesReceiver(
	ws *websocket.Conn,
	cfg component.Config,
) {
	wrapper, err := newTracesReceiverWrapper(h.logger, ws, cfg, h.receiverFactory)
	if err != nil {
		sendErr(ws, h.logger, "failed to create traces receiver wrapper", err)
		return
	}
	h.pipeline.connectTracesReceiverWrapper(wrapper)
	err = wrapper.Start(context.Background(), componenttest.NewNopHost())
	if err != nil {
		sendErr(ws, h.logger, "failed to start traces receiver", err)
		return
	}
	wrapper.waitForStopMessage()
	err = wrapper.Shutdown(context.Background())
	if err != nil {
		sendErr(ws, h.logger, "failed to shut down traces receiver", err)
		return
	}
	h.pipeline.disconnectTracesReceiverWrapper()
}
