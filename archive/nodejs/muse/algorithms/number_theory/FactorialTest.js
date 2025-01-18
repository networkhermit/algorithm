import * as TestRunner from '../../util/TestRunner.js';
import * as Factorial from './Factorial.js';

const testFactorial = () => {
  const sample = [
    [0, 1],
    [1, 1],
    [2, 2],
    [3, 6],
    [4, 24],
    [5, 120],
    [6, 720],
    [7, 5040],
    [8, 40320],
    [9, 362880],
    [10, 3_628_800],
    [11, 39_916_800],
    [12, 479_001_600],
  ];

  for (let i = 0, size = sample.length; i < size; i++) {
    if (Factorial.iterativeProcedure(sample[i][0]) !== sample[i][1]) {
      return false;
    }
  }

  for (let i = 0, size = sample.length; i < size; i++) {
    if (Factorial.recursiveProcedure(sample[i][0]) !== sample[i][1]) {
      return false;
    }
  }

  return true;
};

const main = () => {
  TestRunner.pick(testFactorial);
};

main();
