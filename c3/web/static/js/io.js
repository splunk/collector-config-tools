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

class ComponentFetcher {

  fetchComponents(callback) {
    fetch('/components').then(resp => {
      if (resp.ok) {
        resp.json().then(
          components => {
            callback(components);
          }
        );
      } else {
        console.log('error getting component info');
      }
    });
  }

}

class ConfigMetadataFetcher {

  fetchCfgSchema(componentType, componentName, onMetadataReceived) {
    fetch('/cfg-metadata/' + componentType + '/' + componentName).then(
      resp => {
        if (resp.ok) {
          resp.json().then(
            metadata => {
              flagUnrenderableFields(metadata);
              onMetadataReceived(metadata);
            }
          );
        } else {
          alert("Error getting component schema");
        }
      }
    );
  }

}

function flagUnrenderableFields(cfgSchema) {
  if ((cfgSchema['kind'] === 'struct' || cfgSchema['kind'] === 'ptr') && cfgSchema['fields'] === null) {
    cfgSchema['_unrenderable'] = true;
    return true;
  }
  let unrenderable = false;
  if (cfgSchema['fields'] !== undefined) {
    unrenderable = true;
    cfgSchema['fields'].forEach(field => {
      const fieldUnrenderable = flagUnrenderableFields(field);
      if (!fieldUnrenderable) {
        unrenderable = false;
      }
    });
  }
  if (unrenderable) {
    cfgSchema['_unrenderable'] = true;
  }
  return unrenderable;
}

class YamlConverter {

  jsonToYaml(json, callback) {
    fetch(new Request('/json-to-yaml', {
      method: 'POST',
      body: json
    })).then(
      resp => {
        if (resp.ok) {
          resp.text().then(
            yaml => callback(yaml)
          );
        } else {
          console.log('Error converting JSON to YAML');
        }
      }
    );
  }

  yamlToJson(yaml, callback) {
    fetch(new Request('/yaml-to-json', {
      method: 'POST',
      body: yaml
    })).then(
      resp => {
        if (resp.ok) {
          resp.text().then(
            json => callback(json)
          );
        } else {
          console.log('Error converting YAML to JSON');
        }
      }
    );
  }

}

class Stitcher {

  stitch(collectorYaml, componentCfg, componentGroup, componentName, callback) {
    let req = {
      componentGroup: componentGroup,
      componentName: componentName,
      collectorYaml: collectorYaml,
      componentCfg: componentCfg
    };
    fetch(new Request('/stitch', {
      method: 'POST',
      body: JSON.stringify(req)
    })).then(
      resp => {
        if (resp.ok) {
          resp.text().then(
            yaml => callback(yaml)
          );
        } else {
          console.log('Error converting JSON to YAML');
        }
      }
    );
  }

}

class PipelineCreator {

  createPipeline(pipelineType, collectorYaml, callback) {
    const req = {
      collectorYaml: collectorYaml,
      pipelineType: pipelineType
    };
    fetch(new Request('/create-pipeline', {
      method: 'POST',
      body: JSON.stringify(req)
    })).then(
      resp => {
        if (resp.ok) {
          resp.text().then(
            updatedYaml => callback(updatedYaml)
          );
        } else {
          console.log('Error creating pipeline yaml');
        }
      }
    )
  }

}
