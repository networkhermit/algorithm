'use strict'

const TestRunner = require('muse/util/TestRunner')

const testPick = () => {
  for (let i = 0; i < 10; i++) {
    TestRunner.pick(() => (i & 1) !== 0)
  }
}

if (module === require.main) {
  testPick()
}
