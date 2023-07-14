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

class ComponentSelectorController {

  constructor(sb, eventBus, componentGroup) {
    this.rootView = new ComponentSelectorView(
      sb,
      componentGroup,
      (name, userInputs) => eventBus.emit('component-selected', componentGroup, name, userInputs)
    );
    this.rootView.addOption('------');
    const jsonKey = capitalizeFirstLetter(componentGroup) + 's';
    eventBus.on('components-received', components => this.setOptions(components[jsonKey]));
    eventBus.on('yaml-to-json-received', cfg => this.setConfiguredComponents(cfg[componentGroup + 's']));
    eventBus.on('component-form-submitted', () => this.reset());
  }

  getRootView() {
    return this.rootView;
  }

  setOptions(components) {
    components.forEach(component => {
      this.rootView.addOption(component['Name']);
    });
  }

  reset() {
    this.rootView.selectIndex(0);
  }

  setConfiguredComponents(components) {
    this.rootView.setConfiguredComponents(components);
  }

}

class ComponentSelectorView extends View {

  constructor(sb, componentGroup, onSelected) {
    super();
    this.sb = sb;
    const componentGroupTitle = capitalizeFirstLetter(componentGroup) + 's';
    this.onSelected = onSelected;

    const fieldsetWidget = new FieldsetWidget(componentGroupTitle);
    sb.styleComponentSelectorFieldsetView(fieldsetWidget);
    this.appendView(fieldsetWidget);

    this.selectWidget = new SelectWidget();
    sb.styleSelectWidget(this.selectWidget);
    this.selectWidget.onSelected(() => {
      onSelected(this.selectWidget.getValue());
    });
    fieldsetWidget.appendView(this.selectWidget);
    this.configuredComponentsPanelView = new DivWidget('configuredComponentsPanelView');
    sb.styleConfiguredComponentsPanelView(this.configuredComponentsPanelView);
    fieldsetWidget.appendView(this.configuredComponentsPanelView);
  }

  selectIndex(i) {
    this.selectWidget.selectIndex(i);
  }

  addOption(option) {
    this.selectWidget.addOption(option, option);
  }

  setConfiguredComponents(components) {
    this.configuredComponentsPanelView.removeChildren();
    for (const c in components) {
      // fixme: this view is creating a controller
      const configuredComponentController = new ConfiguredComponentController(this.sb, c, components[c], this.onSelected);
      this.configuredComponentsPanelView.appendView(configuredComponentController.getRootView());
    }
  }

}
