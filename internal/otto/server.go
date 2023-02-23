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

	mux.Handle("/cfg-metadata/", cfgschemaHandler{
		logger: logger,
	})

	mux.Handle("/json-to-yaml", jsonToYAMLHandler{
		logger: logger,
	})

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
