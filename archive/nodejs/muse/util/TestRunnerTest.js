import * as TestRunner from './TestRunner.js';

const testPick = () => {
  for (let i = 0; i < 10; i++) {
    TestRunner.pick(() => (i & 1) !== 0);
  }
};

const main = () => {
  testPick();
};

main();
