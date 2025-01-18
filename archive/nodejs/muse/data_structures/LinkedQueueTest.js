import * as TestRunner from '../util/TestRunner.js';
import * as LinkedQueue from './LinkedQueue.js';

const testLinkedQueue = () => {
  const size = 8192;

  const queue = new LinkedQueue.LinkedQueue();

  for (let i = 1; i <= size; i++) {
    queue.enqueue(i);
  }

  if (queue.size() !== size) {
    return false;
  }

  for (let i = 1; i <= size; i++) {
    if (queue.peek() !== i) {
      return false;
    }
    queue.dequeue();
  }

  return queue.isEmpty();
};

const main = () => {
  TestRunner.pick(testLinkedQueue);
};

main();
