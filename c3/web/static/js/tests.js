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
  testEventBus();
}

function testEventBus() {

  class X {

    constructor(eventBus) {
      this.eventBus = eventBus;
      this.val = 0;
      this.callback = () => this.val++;
      eventBus.on('event', this.callback);
    }

    unregister() {
      this.eventBus.off('event', this.callback);
    }

  }

  const eb = new EventBus();
  const x = new X(eb);
  assertExpectedEqActual(0, x.val);
  eb.emit('event');
  assertExpectedEqActual(1, x.val);
  x.unregister();
  eb.emit('event');
  assertExpectedEqActual(1, x.val);
}
