import * as TestRunner from '../util/TestRunner.js';
import { LinkedStack } from './LinkedStack.js';

const testLinkedStack = () => {
  const size = 8192;

  const stack = new LinkedStack();

  for (let i = 1; i <= size; i++) {
    stack.push(i);
  }

  if (stack.size() !== size) {
    return false;
  }

  for (let i = size; i > 0; i--) {
    if (stack.peek() !== i) {
      return false;
    }
    stack.pop();
  }

  return stack.isEmpty();
};

const main = () => {
  TestRunner.pick(testLinkedStack);
};

main();
