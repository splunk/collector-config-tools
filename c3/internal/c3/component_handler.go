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
	"log"
	"net/http"
)

type componentHandler struct {
	logger *log.Logger
	jsn    []byte
}

func (h componentHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_, err := resp.Write(h.jsn)
	if err != nil {
		h.logger.Printf("componentHandler: ServeHTTP: error writing response: %v", err)
	}
}
