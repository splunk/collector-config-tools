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


class CreatePipelinePanelContoller {

  constructor(sb, eventBus) {
    this.rootView = new CreatePipelinePanelView(sb);

    const metricsButton = new CreatePipelineButtonView('metrics', sb);
    metricsButton.onClick(() => eventBus.emit('metrics-button-clicked'));
    this.rootView.appendView(metricsButton);

    const tracesButton = new CreatePipelineButtonView('traces', sb);
    this.rootView.appendView(tracesButton);

    const logsButton = new CreatePipelineButtonView('logs', sb);
    this.rootView.appendView(logsButton);
  }

  getRootView() {
    return this.rootView;
  }

}

class CreatePipelinePanelView extends FieldsetWidget {

  constructor(sb) {
    super('Create a Pipeline');
    sb.stylePipelinePanelView(this);
  }

}

class CreatePipelineButtonView extends ButtonWidget {

  constructor(componentGroup, sb) {
    super(capitalizeFirstLetter(componentGroup));
    sb.styleButtonWidget(this);
  }

}
