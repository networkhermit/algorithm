import * as SequenceBuilder from '../../util/SequenceBuilder.js';
import * as Sequences from '../../util/Sequences.js';
import * as TestRunner from '../../util/TestRunner.js';
import * as BubbleSort from './BubbleSort.js';

const testBubbleSort = () => {
  const size = 32768;

  const arr = new Array(size);
  SequenceBuilder.packRandom(arr);

  const checksum = Sequences.parityChecksum(arr);

  BubbleSort.sort(arr);

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false;
  }

  return Sequences.isSorted(arr);
};

const main = () => {
  TestRunner.pick(testBubbleSort);
};

main();
