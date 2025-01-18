import * as TestRunner from '../util/TestRunner.js';
import { ArrayList } from './ArrayList.js';

const testArrayList = () => {
  const size = 8192;

  const list = new ArrayList(0);

  for (let i = 1; i <= size; i++) {
    list.append(i);
  }

  list.shrink();

  if (list.size() !== size) {
    return false;
  }

  if (list.capacity() !== size) {
    return false;
  }

  for (let i = 0; i < size; i++) {
    if (list.get(i) !== i + 1) {
      return false;
    }
  }

  for (let i = 0; i < size; i++) {
    list.set(i, size - i);
  }

  for (let i = 0; i < size; i++) {
    if (list.get(i) !== size - i) {
      return false;
    }
  }

  for (let i = 0, j = size - 1; i < j; i++, j--) {
    const x = list.get(i);
    const y = list.get(j);

    list.remove(i);
    list.insert(i, y);
    list.remove(j);
    list.insert(j, x);
  }

  for (let i = 0; i < size; i++) {
    if (list.get(i) !== i + 1) {
      return false;
    }
  }

  for (let i = size; i >= 1; i--) {
    if (list.back() !== i) {
      return false;
    }
    list.eject();
  }

  list.shrink();

  if (!list.isEmpty()) {
    return false;
  }

  return list.capacity() === 0;
};

const main = () => {
  TestRunner.pick(testArrayList);
};

main();
