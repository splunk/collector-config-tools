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

class YamlPanelController {

  constructor(sb, eventBus, pipelineCreator) {
    this.yamlView = new YamlPanelView(sb, () => {
      eventBus.emit('yaml-updated', this.getYaml());
    });
    eventBus.on('metrics-button-clicked', () => {
      pipelineCreator.createPipeline('metrics', this.getYaml(), updatedYaml => {
        this.setYaml(updatedYaml);
      });
    });
  }

  getRootView() {
    return this.yamlView;
  }

  setYaml(yaml) {
    this.yamlView.setYaml(yaml);
  }

  getYaml() {
    return this.yamlView.getYaml();
  }

}

class YamlPanelView extends View {

  constructor(sb, onApplyClicked) {
    super();
    this.addClass('YamlPanelView');
    sb.styleYamlPanelView(this);
    this.ta = new TextareaWidget(sb);
    sb.styleYamlPanelTextAreaWidget(this.ta);
    this.appendView(this.ta);
    const buttonWidget = new ButtonWidget('Apply');
    sb.styleButtonWidget(buttonWidget);
    buttonWidget.addEventListener('click', onApplyClicked);
    this.appendView(buttonWidget);
  }

  setYaml(yaml) {
    this.ta.setText(yaml);
  }

  getYaml() {
    return this.ta.getText();
  }

}
