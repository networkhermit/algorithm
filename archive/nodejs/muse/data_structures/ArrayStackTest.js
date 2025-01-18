import * as TestRunner from '../util/TestRunner.js';
import * as ArrayStack from './ArrayStack.js';

const testArrayStack = () => {
  const size = 8192;

  const stack = new ArrayStack.ArrayStack(0);

  for (let i = 1; i <= size; i++) {
    stack.push(i);
  }

  stack.shrink();

  if (stack.size() !== size) {
    return false;
  }

  if (stack.capacity() !== size) {
    return false;
  }

  for (let i = size; i > 0; i--) {
    if (stack.peek() !== i) {
      return false;
    }
    stack.pop();
  }

  stack.shrink();

  if (!stack.isEmpty()) {
    return false;
  }

  return stack.capacity() === 0;
};

const main = () => {
  TestRunner.pick(testArrayStack);
};

main();
