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

class SidebarController {

  constructor(sb, eventBus, componentFetcher, yamlConverter) {
    this.eventBus = eventBus;
    this.sidebarView = new SidebarView(sb);
    this.componentFetcher = componentFetcher;
    this.yamlConverter = yamlConverter;

    eventBus.on('yaml-updated', yaml => this.yamlUpdated(yaml));
  }

  getRootView() {
    return this.sidebarView;
  }

  fetchComponents() {
    this.componentFetcher.fetchComponents(components => {
      this.eventBus.emit('components-received', components);
    });
  }

  appendView(v) {
    this.sidebarView.appendView(v);
  }

  yamlUpdated(yaml) {
    this.yamlConverter.yamlToJson(yaml, json => {
      const cfg = JSON.parse(json)
      this.eventBus.emit('yaml-to-json-received', cfg);
    });
  }

}

class SidebarView extends View {

  constructor(sb) {
    super();
    this.addClass('SidebarView');
    sb.styleSidebarView(this);
    const labelWidget = new LabelWidget('C3');
    sb.styleSidebarLabelWidget(labelWidget);
    this.appendView(labelWidget);
  }

}
