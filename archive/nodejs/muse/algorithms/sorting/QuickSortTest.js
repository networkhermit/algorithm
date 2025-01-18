import * as SequenceBuilder from '../../util/SequenceBuilder.js';
import * as Sequences from '../../util/Sequences.js';
import * as TestRunner from '../../util/TestRunner.js';
import * as QuickSort from './QuickSort.js';

const testQuickSort = () => {
  const size = 32768;

  const arr = new Array(size);
  SequenceBuilder.packRandom(arr);

  const checksum = Sequences.parityChecksum(arr);

  QuickSort.sort(arr);

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false;
  }

  return Sequences.isSorted(arr);
};

const main = () => {
  TestRunner.pick(testQuickSort);
};

main();
