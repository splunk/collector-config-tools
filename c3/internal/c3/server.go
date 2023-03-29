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
	"log"
	"net/http"

	"go.opentelemetry.io/collector/otelcol"
)

func Server(logger *log.Logger, addr string, factories otelcol.Factories) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web/static")))

	componentTypes := factoriesToComponentTypes(factories)
	jsn, err := json.Marshal(componentTypes)
	if err != nil {
		logger.Printf("Server: error getting components: %v", err)
		return
	}

	mux.Handle("/components", componentHandler{
		logger: logger,
		jsn:    jsn,
	})

	mux.Handle("/cfg-metadata/", cfgschemaHandler{
		logger:         logger,
		componentTypes: componentTypes,
	})

	mux.Handle("/json-to-yaml", jsonToYAMLHandler{
		logger: logger,
	})

	mux.Handle("/yaml-to-json", yamlToJSONHandler{
		logger: logger,
	})

	mux.Handle("/stitch", stitchHandler{
		logger: logger,
	})

	mux.Handle("/create-pipeline", createPipelineHandler{
		logger: logger,
	})

	svr := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	println("c3 started", addr)
	err = svr.ListenAndServe()
	if err != nil {
		logger.Fatalf("http serve error: %v", err)
	}
}
