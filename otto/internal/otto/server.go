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
	"log"
	"net/http"

	"go.opentelemetry.io/collector/otelcol"
)

func Server(logger *log.Logger, addr string, factories otelcol.Factories) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web/static")))

	mux.Handle("/components", componentHandler{
		logger:    logger,
		factories: factories,
	})

	ottoPipeline := &pipeline{
		factories: factories,
	}
	mux.Handle("/cfg-metadata/", cfgschemaHandler{
		logger:   logger,
		pipeline: ottoPipeline,
	})

	mux.Handle("/jsonToYAML", jsonToYAMLHandler{
		logger: logger,
	})

	wsHandlers := map[string]wsHandler{}
	registerReceiverHandlers(logger, factories, wsHandlers, ottoPipeline)
	registerProcessorHandlers(logger, factories, wsHandlers, ottoPipeline)
	registerExporterHandlers(logger, factories, wsHandlers, ottoPipeline)
	mux.Handle("/ws/", httpWsHandler{handlers: wsHandlers})

	svr := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	println("otto started")
	err := svr.ListenAndServe()
	if err != nil {
		logger.Fatalf("http serve error: %v", err)
	}
}

func registerReceiverHandlers(logger *log.Logger, factories otelcol.Factories, handlers map[string]wsHandler, ppln *pipeline) {
	for componentName, factory := range factories.Receivers {
		const componentType = "receiver"
		path := "/ws/" + componentType + "/" + componentName.String()
		handlers[path] = receiverSocketHandler{
			logger:          logger,
			pipeline:        ppln,
			receiverFactory: factory,
		}
	}
}

func registerProcessorHandlers(logger *log.Logger, factories otelcol.Factories, handlers map[string]wsHandler, ppln *pipeline) {
	for componentName, factory := range factories.Processors {
		const componentType = "processor"
		path := "/ws/" + componentType + "/" + componentName.String()
		handlers[path] = processorSocketHandler{
			logger:           logger,
			pipeline:         ppln,
			processorFactory: factory,
		}
	}
}

func registerExporterHandlers(logger *log.Logger, factories otelcol.Factories, handlers map[string]wsHandler, ppln *pipeline) {
	for componentName, factory := range factories.Exporters {
		const componentType = "exporter"
		path := "/ws/" + componentType + "/" + componentName.String()
		handlers[path] = exporterSocketHandler{
			logger:          logger,
			pipeline:        ppln,
			exporterFactory: factory,
		}
	}
}
