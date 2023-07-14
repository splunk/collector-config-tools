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

window.onload = main;

function main() {
  const eventBus = new EventBus();
  const sb = new StyleBundle();

  const mainView = new MainView(document, sb)

  const sidebarController = new SidebarController(sb, eventBus, new ComponentFetcher(), new YamlConverter());

  const receiverSelectorController = new ComponentSelectorController(sb, eventBus, 'receiver');
  sidebarController.appendView(receiverSelectorController.getRootView());

  const processorSelectorController = new ComponentSelectorController(sb, eventBus, 'processor');
  sidebarController.appendView(processorSelectorController.getRootView());

  const exporterSelectorController = new ComponentSelectorController(sb, eventBus, 'exporter');
  sidebarController.appendView(exporterSelectorController.getRootView());

  const createPipelinePanelContoller = new CreatePipelinePanelContoller(sb, eventBus);
  sidebarController.appendView(createPipelinePanelContoller.getRootView());

  mainView.appendView(sidebarController.getRootView());

  const configPanelController = new ConfigPanelController(
    sb,
    eventBus,
    new ConfigMetadataFetcher(),
    new Stitcher(),
    new PipelineCreator()
  );
  mainView.appendView(configPanelController.getRootView());

  sidebarController.fetchComponents();
}
