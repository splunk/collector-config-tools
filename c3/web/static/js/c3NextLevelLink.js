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


class C3NextLevelLinkController {

  constructor(cfgSchema, parentKey, parentController, parentView, sb, docStr) {
    this.parentKey = parentKey;
    this.parentController = parentController;
    this.parentView = parentView;
    this.cfgSchema = cfgSchema;
    this.sb = sb;
    this.nextLevelLinkView = new C3NextLevelLinkView(
      cfgSchema["name"],
      docStr,
      sb,
      () => this.handleClick(),
    );
  }

  getView() {
    return this.nextLevelLinkView;
  }

  highlightLink() {
    this.nextLevelLinkView.highlight();
  }

  unhighlightLink() {
    this.nextLevelLinkView.unhighlight();
  }

  markLinkAsClicked() {
    this.nextLevelLinkView.markAsClicked();
  }

  handleClick() {
    this.highlightLink()
    this.parentController.nextLevelClicked();
    const userInputs = this.parentController.getUserInputsForKey(this.parentKey)
    switch (this.cfgSchema['kind']) {
      case 'slice':
        new C3FieldsetSliceFormController(
          this.sb,
          this.cfgSchema,
          this.parentView,
          userInputs => this.childFormSubmitted(userInputs)
        );
        break;
      case 'map':
        new C3FieldsetMapFormController(
          this.sb,
          this.cfgSchema,
          this.parentView,
          userInputs => this.childFormSubmitted(userInputs)
        );
        break;
      default:
        const c = new C3FieldsetFormController(
          this.cfgSchema,
          this.parentView,
          this.sb,
          userInputs => this.childFormSubmitted(userInputs),
          userInputs,
          true
        );
        c.render();
    }
  }

  childFormSubmitted(userInputs) {
    this.parentController.captureChildSubmission(userInputs, this.parentKey);
    this.unhighlightLink();
    this.markLinkAsClicked();
  }

}

class C3NextLevelLinkView extends View {

  constructor(labelStr, docStr, sb, onclick) {
    super();
    sb.styleNextLevelLinkView(this);
    this.addClass('C3NextLevelLinkView');
    this.appendView(new FieldLabelView(sb, labelStr, docStr));
    const inputWrapperView = new DivWidget('input');
    this.appendView(inputWrapperView);
    this.linkWidget = new LinkWidget('[configure]', onclick, 'next-level');
    sb.styleNextLevelLinkWidget(this.linkWidget);
    inputWrapperView.appendView(this.linkWidget);
  }

  highlight() {
    this.linkWidget.addClass('highlighted');
  }

  unhighlight() {
    this.linkWidget.removeClass('highlighted');
  }

  markAsClicked() {
    this.linkWidget.addClass('clicked');
  }

  disable() {
    this.linkWidget.addClass('disabled');
  }

  enable() {
    this.linkWidget.removeClass('disabled');
  }

}
