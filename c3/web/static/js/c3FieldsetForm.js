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

class C3FieldsetFormController {

  constructor(cfgSchema, parentView, sb, onSubmit, userInputs, isSubField) {
    this.cfgSchema = cfgSchema;
    this.parentView = parentView;
    this.sb = sb;
    this.childSchemasByName = {};
    this.userInputs = userInputs === undefined ? {} : userInputs;
    this.onSubmit = onSubmit;
    this.isSubField = isSubField;
    const name = cfgSchemaToFieldsetName(cfgSchema);
    this.rootView = new C3FieldsetFormView(name, sb, cfgSchema['doc']);
    this.rootView.onApplyButtonClick(() => this.submitForm());
    parentView.appendView(this.rootView);
  }

  render() {
    this.cfgSchema['fields'].forEach(cfgSchema => {
      const fieldName = cfgSchema['name'];
      this.childSchemasByName[fieldName] = cfgSchema;
      const userInput = this.userInputs === undefined ? undefined : this.userInputs[fieldName];
      this.renderField(cfgSchema, userInput);
    });
  }

  renderField(cfgSchema, userInput) {
    const kind = cfgSchema['kind'];
    if (kind === 'interface' || kind === 'func') {
      return;
    }
    const fieldName = cfgSchema['name'];
    const docStr = cfgSchemaToDocStr(cfgSchema);
    if (kind === 'struct' || kind === 'ptr' || kind === 'slice' || kind === 'map') {
      this.renderCompoundField(cfgSchema, fieldName, docStr);
    } else {
      this.renderInlineField(cfgSchema, kind, userInput, fieldName, docStr);
    }
  }

  renderCompoundField(cfgSchema, key, docStr) {
    const nextLevelLinkController = new C3NextLevelLinkController(
      cfgSchema,
      key,
      this,
      this.parentView,
      this.sb,
      docStr
    );
    const nextLevelLinkView = nextLevelLinkController.getView();
    this.rootView.registerLinkView(nextLevelLinkView);
    this.rootView.appendToFormView(nextLevelLinkView);
  }

  nextLevelClicked() {
    this.rootView.disableInputs();
  }

  captureChildSubmission(childInputs, userInputKey) {
    this.userInputs[userInputKey] = childInputs;
    this.rootView.enableInputs();
  }

  getUserInputsForKey(key) {
    return this.userInputs[key];
  }

  renderInlineField(cfgSchema, kind, userInput, fieldName, docStr) {
    const defaultVal = cfgSchema['default'];
    let widget;
    switch (kind) {
      case "bool":
        widget = new C3BoolSelectView(cfgSchema.name, defaultVal);
        this.sb.styleSelectWidget(widget);
        break;
      default:
        const placeholder = defaultVal != null ? defaultVal : undefined;
        widget = new TextInputWidget(cfgSchema.name, placeholder, userInput);
        this.sb.styleInputWidget(widget);
        break;
    }
    const fieldView = new C3FieldView(this.sb, fieldName, docStr);
    fieldView.appendInputWidget(widget);
    this.rootView.appendToFormView(fieldView);
  }

  submitForm() {
    this.rootView.forEachFormElement(el => {
      if (el.value !== '') {
        const childSchema = this.childSchemasByName[el.name];
        this.userInputs[el.name] = convertUserInput(el.value, childSchema['kind'], childSchema['type']);
      }
    });
    this.onSubmit(this.userInputs);
    if (this.isSubField) {
      this.removeView();
    }
  }

  removeView() {
    this.parentView.removeView(this.rootView);
  }

}

function convertUserInput(userInput, kind, type) {
  switch (kind) {
    case "int":
      return parseInt(userInput);
    case "int64":
      return type === 'time.Duration' ? userInput : parseInt(userInput);
    case "bool":
      return userInput === 'true';
    default:
      return userInput;
  }
}

function cfgSchemaToFieldsetName(cfgSchema) {
  let out = cfgSchema['name'] || cfgSchema['type'];
  if (out.charAt(0) === '*') {
    out = out.substring(1);
  }
  return out;
}

function cfgSchemaToDocStr(cfgSchema) {
  let out = cfgSchema['doc'];
  const kind = cfgSchema['kind'];
  if (kind != null) {
    out += '\nKind: ' + kind;
  }
  const typ = cfgSchema['type'];
  if (typ !== undefined && typ.length > 0) {
    if (out.length > 0) {
      out += '\n';
    }
    out += 'Type: ' + typ;
  }
  return out;
}

class C3FieldsetFormView extends View {

  constructor(name, sb, doc) {
    super();

    this.addClass('C3FieldsetFormView');
    this.linkViews = [];
    this.formWidget = new FormWidget();
    this.btnWidget = new ButtonWidget('Apply');
    sb.styleButtonWidget(this.btnWidget);
    const fieldset = new FieldsetWidget(name);
    sb.styleC3FieldsetFormView(fieldset);
    fieldset.appendView(this.formWidget);
    fieldset.appendView(this.btnWidget);
    this.appendView(fieldset);
  }

  appendToFormView(view) {
    this.formWidget.appendView(view);
  }

  registerLinkView(linkView) {
    this.linkViews.push(linkView);
  }

  onApplyButtonClick(f) {
    this.btnWidget.addEventListener('click', f);
  }

  forEachFormElement(f) {
    this.formWidget.forEachFormElement(f);
  }

  disableInputs() {
    this.btnWidget.setInvisible();
    this.linkViews.forEach(lv => lv.disable())
  }

  enableInputs() {
    this.btnWidget.setVisible();
    this.linkViews.forEach(lv => lv.enable())
  }

}

class FieldLabelView extends View {

  constructor(sb, labelStr, docStr) {
    super();
    this.addClass('FieldLabelView');
    sb.styleFieldLabelView(this);
    const inputLabelWidget = new LabelWidget(labelStr);
    this.appendView(inputLabelWidget);
    this.appendView(new ToolTipView(sb, docStr));
  }

}

class ToolTipView extends LabelWidget {

  constructor(sb, docStr) {
    super('\u{24D8}');
    this.addClass('ToolTipView');
    sb.styleInfoLabelWidget(this);
    this.setAttribute('title', docStr);
  }

}
