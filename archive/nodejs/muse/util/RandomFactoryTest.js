import * as RandomFactory from './RandomFactory.js';
import * as TestRunner from './TestRunner.js';

const testIntegerN = () => {
  let value = 0;
  for (let i = 0; i < 8192; i++) {
    if (RandomFactory.genIntN(0, 0) !== 0) {
      return false;
    }

    if (RandomFactory.genIntN(1, 1) !== 1) {
      return false;
    }

    value = RandomFactory.genIntN(0, 1);
    if (value < 0 || value > 1) {
      return false;
    }

    value = RandomFactory.genIntN(100, 10000);
    if (RandomFactory.genIntN(value, value) !== value) {
      return false;
    }
    if (value < 100 || value > 10000) {
      return false;
    }
  }

  return true;
};

const testGenEven = () => {
  for (let i = 0; i < 8192; i++) {
    if ((RandomFactory.genEven() & 1) !== 0) {
      return false;
    }
  }

  return true;
};

const testGenOdd = () => {
  for (let i = 0; i < 8192; i++) {
    if ((RandomFactory.genOdd() & 1) === 0) {
      return false;
    }
  }

  return true;
};

const main = () => {
  TestRunner.pick(testIntegerN);

  TestRunner.pick(testGenEven);

  TestRunner.pick(testGenOdd);
};

main();
