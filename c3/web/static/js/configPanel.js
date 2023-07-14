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

class ConfigPanelController {

  constructor(sb, eventBus, configMetadataFetcher, stitcher, pipelineCreator) {
    this.sb = sb;
    this.eventBus = eventBus;
    this.configMetadataFetcher = configMetadataFetcher;
    this.stitcher = stitcher;

    this.rootView = new DivWidget('ConfigPanelController-rootView');
    sb.styleComponentPanelRootView(this.rootView);

    const tabBarContainerView = new TabBarContainerView();
    sb.styleTabBarContainerView(tabBarContainerView);
    this.rootView.appendView(tabBarContainerView);

    const tabPanelContainerView = new TabPanelContainerView(sb);
    this.rootView.appendView(tabPanelContainerView);

    const deselector = v => sb.styleDeselectedTab(v);
    const selector = v => sb.styleSelectedTab(v);
    this.tabController = new TabController(tabBarContainerView, tabPanelContainerView, deselector, selector);

    const tabWrapperStyler = v => sb.styleTabContentWrapperView(v);

    this.yamlPanelController = new YamlPanelController(sb, eventBus, pipelineCreator);
    this.tabController.addTab('YAML', this.yamlPanelController.getRootView(), tabWrapperStyler);
    this.tabController.select(0);

    eventBus.on('component-selected', (group, name, userInputs) => this.componentSelected(group, name, userInputs));
  }

  addComponentTab(tabName) {
    this.formPanelContainerView = new DivWidget('formPanelContainerView');
    this.sb.styleFormPanelContainerView(this.formPanelContainerView);
    this.tabController.addTab(tabName, this.formPanelContainerView, tabView => this.styleTab(tabView));
  }

  removeComponentTab() {
    this.tabController.removeTab(1);
  }

  styleTab(tabView) {
    this.sb.styleTabContentWrapperView(tabView);
  }

  getRootView() {
    return this.rootView;
  }

  getYaml() {
    return this.yamlPanelController.getYaml();
  }

  componentSelected(componentGroup, componentName, userInputs) {
    this.addComponentTab(componentGroup + '/' + componentName);
    if (this.rootFieldsetStructFormController !== undefined) {
      this.rootFieldsetStructFormController.removeView();
      this.rootFieldsetStructFormController = undefined;
    }
    this.configMetadataFetcher.fetchCfgSchema(componentGroup, componentName, cfgSchema => {
      this.renderComponentFieldsetView(cfgSchema, componentGroup, componentName, userInputs);
      this.tabController.select(1);
    });
  }

  renderComponentFieldsetView(cfgSchema, componentGroup, componentName, userInputs) {
    this.rootFieldsetStructFormController = new C3FieldsetFormController(
      cfgSchema,
      this.formPanelContainerView,
      this.sb,
      inputs => {
        this.userInputsReceived(inputs, componentGroup, componentName);
        this.eventBus.emit('component-form-submitted');
      },
      userInputs
    );
    this.rootFieldsetStructFormController.render();
  }

  userInputsReceived(componentCfg, componentGroup, componentName) {
    const currYaml = this.yamlPanelController.getYaml();
    this.stitcher.stitch(currYaml, componentCfg, componentGroup, componentName, yaml => {
      this.yamlPanelController.setYaml(yaml);
      this.eventBus.emit('yaml-updated', yaml);
      this.tabController.select(0);
      this.removeComponentTab();
    });
    this.rootFieldsetStructFormController.removeView();
    this.rootFieldsetStructFormController = undefined;
  }

}

class TabBarContainerView extends View {

  constructor() {
    super();
    this.addClass('TabBarContainerView');
  }

}

class TabPanelContainerView extends View {

  constructor(sb) {
    super();
    this.addClass('TabPanelContainerView');
    sb.styleTabPanelContainerView(this);
  }

}
