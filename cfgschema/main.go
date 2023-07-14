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

package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/splunk/collector-config-tools/cfgschema/lib/configschema"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

func run() error {
	sourceDir, outputDir := getFlags()
	c, err := components()
	if err != nil {
		return err
	}
	return configschema.GenerateYAMLFiles(c, sourceDir, outputDir, "github.com/splunk/collector-config-tools/cfgschema")
}

func getFlags() (string, string) {
	sourceDir := flag.String("s", filepath.Join("."), "")
	outputDir := flag.String("o", filepath.Join("..", "cfg-metadata"), "output dir")
	flag.Parse()
	return *sourceDir, *outputDir
}
