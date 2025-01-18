import * as TestRunner from '../util/TestRunner.js';
import * as ArrayQueue from './ArrayQueue.js';

const testArrayQueue = () => {
  const size = 8192;

  const queue = new ArrayQueue.ArrayQueue(0);

  for (let i = 1; i <= size; i++) {
    queue.enqueue(i);
  }

  queue.shrink();

  if (queue.size() !== size) {
    return false;
  }

  if (queue.capacity() !== size) {
    return false;
  }

  for (let i = 1; i <= size; i++) {
    if (queue.peek() !== i) {
      return false;
    }
    queue.dequeue();
  }

  queue.shrink();

  if (!queue.isEmpty()) {
    return false;
  }

  return queue.capacity() === 0;
};

const main = () => {
  TestRunner.pick(testArrayQueue);
};

main();
