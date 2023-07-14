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

class C3FieldView extends View {

  constructor(sb, labelStr, docStr) {
    super();
    sb.styleFieldView(this);
    this.addClass('C3FieldView');

    this.appendView(new FieldLabelView(sb, labelStr, docStr));

    this.inputWrapperView = new DivWidget('input');
    this.appendView(this.inputWrapperView);
  }

  appendInputWidget(w) {
    this.inputWrapperView.appendView(w);
  }

}
