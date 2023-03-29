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

class StyleBundle {

  constructor() {
    this.darkBgColor = 'rgb(16, 18, 22)';
    this.sidebarBgColor = 'rgb(14, 17, 22)';
    this.textColor = 'rgb(202, 209, 216)';
    this.buttonTextColor = 'rgb(105, 166, 248)';
    this.textPillColor = 'rgb(82, 134, 244)';
    this.textPillBgColor = 'rgb(21, 29, 45)';
    this.inputWidgetBgColor = 'rgb(1, 5, 11)';
    this.buttonBorderColor = 'rgb(49, 80, 52)';
    this.borderColor = 'rgb(49, 54, 60)';
  }

  styleTestWrapper(v) {
    v.setDisplay('flex');
    v.setJustifyContent('center');
    v.setAlignItems('center');
    v.setHeight('100%');
    v.setWidth('100%');
  }

  styleMainView(v) {
    v.setBackgoundColor(this.darkBgColor);
    v.setColor(this.textColor);
    v.setMargin(0);
    v.setDisplay('flex');
    v.setHeight('100%');
    v.setFontFamily('sans-serif');
  }

  styleSidebarView(v) {
    this.padTopLevelPanel(v);
    v.setWidth('400px');
    v.setBackgoundColor(this.sidebarBgColor);
    v.setDisplay('flex');
    v.setFlexDirection('column');
    v.setGap('12px');
  }

  styleSidebarLabelWidget(v) {
    v.setFontWeight('bold');
    v.setFontSize('18px');
    v.setMargin('0 0 0 2px');
  }

  styleComponentPanelRootView(v) {
    v.setFlexGrow('1');
    v.setPadding('24px');
  }

  styleTabBarContainerView(v) {
    v.setDisplay('flex');
    v.setGap('12px');
    v.setPadding('0 0 12px 4px');
  }

  styleDeselectedTab(v) {
    this.styleTab(v);
    v.setBorderColor('rgb(0, 0, 0)');
  }

  styleSelectedTab(v) {
    this.styleTab(v);
    v.setBorderColor('rgb(242, 118, 96)');
  }

  styleTab(v) {
    v.setFontSize('18px');
    v.setBorderWidth('0 0 3px 0');
    v.setBorderStyle('solid');
    v.setCursor('pointer');
  }

  styleTabPanelContainerView(v) {
    v.setHeight('96%');
  }

  styleTabContentWrapperView(v) {
    v.setHeight('100%');
  }

  styleYamlPanelView(v) {
    v.setWidth('100%');
    v.setHeight('100%');
  }

  styleYamlPanelTextAreaWidget(v) {
    v.setHeight('75%');
    v.setWidth('95%');
    v.setColor(this.textColor);
    v.setBackgoundColor(this.inputWidgetBgColor);
    v.setFontSize('18px');
    v.setBorderRadius('4px');
    v.setBorderColor(this.borderColor);
    v.setPadding('4px 8px');
  }

  styleC3FieldsetFormView(v) {
    this.styleFieldset(v);
  }

  styleComponentSelectorFieldsetView(v) {
    this.styleFieldset(v);
  }

  stylePipelinePanelView(v) {
    v.setDisplay('flex');
    v.setJustifyContent('space-around');
    this.styleFieldset(v);
  }

  styleFormPanelContainerView(v) {
    v.setDisplay('flex');
    v.setOverflow('auto');
  }

  styleFieldset(v) {
    v.setPadding('12px');
    v.setBorderWidth('1px');
    v.setBorderColor(this.borderColor);
    v.setBorderRadius('4px');
  }

  styleConfiguredComponentsPanelView(v) {
    v.setDisplay('flex');
    v.setGap('4px');
    v.setMarginBottom('-4px'); // naughty
  }

  styleConfiguredComponentView(v) {
    v.setCursor('pointer');
    v.setColor(this.textPillColor);
    v.setBackgoundColor(this.textPillBgColor);
    v.setBorderRadius('16px');
    v.setPadding('2px 12px');
    v.setMargin('8px 0');
    v.setFontSize('1.24em');
    v.setFontFamily('monospace');
  }

  styleLinkWidget(v) {
    this.styleClickable(v);
  }

  styleClickable(v) {
    v.setCursor('pointer');
    v.setTextDecoration('underline');
  }

  styleSelectWidget(v) {
    v.setMargin('4px 0');
    v.setWidth('100%');
    this.styleInputWidget(v);
  }

  styleFieldLabelView(v) {
    v.setDisplay('flex');
    v.setMargin('0 0 0 2px');
    v.setFontFamily('monospace');
    v.setFontSize('1.3em');
  }

  styleInputWidget(v) {
    v.setColor(this.textColor);
    v.setBackgoundColor(this.inputWidgetBgColor);
    v.setBorder('solid rgb(42, 48, 53) 1px');
    v.setBorderRadius('6px');
    v.setPadding('4px');
    v.setFontSize('1.3em');
    v.setFontFamily('monospace');
    v.setMargin('4px 0');
  }

  styleButtonWidget(v) {
    v.setColor(this.buttonTextColor);
    v.setBackgoundColor(this.inputWidgetBgColor);
    v.setFontWeight('bold');
    v.setPadding('6px 12px');
    v.setBorderStyle('solid');
    v.setBorderWidth('1px');
    v.setBorderColor(this.buttonBorderColor);
    v.setBorderRadius('8px');
    v.setFontSize('1.05em');
    v.setCursor('pointer');
  }

  styleFieldView(v) {
    this.padFormField(v);
  }

  styleNextLevelLinkView(v) {
    this.padFormField(v);
  }

  styleNextLevelLinkWidget(v) {
    this.styleLinkWidget(v);
  }

  styleInfoLabelWidget(v) {
    v.setCursor('pointer');
    v.setFontSize('0.8em');
    v.setMargin('0 0 0 4px');
    v.setPosition('relative');
    v.setTop('-1px');
  }

  padFormField(v) {
    v.setMargin('16px 0');
  }

  padTopLevelPanel(v) {
    v.setPadding('24px');
  }

}
