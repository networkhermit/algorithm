package muse.data_structures;

import muse.data_structures.interfaces.IStack;

public class ArrayStack<E> implements IStack<E> {
  private static final int DEFAULT_CAPACITY = 64;

  private E[] data = null;
  private int logicalSize = 0;
  private int physicalSize = DEFAULT_CAPACITY;

  public ArrayStack() {
    this(0);
  }

  @SuppressWarnings("unchecked")
  public ArrayStack(int physicalSize) {
    if (physicalSize > 1) {
      this.physicalSize = physicalSize;
    }
    data = (E[]) new Object[this.physicalSize];
  }

  @Override
  public int size() {
    return logicalSize;
  }

  @Override
  public boolean isEmpty() {
    return logicalSize == 0;
  }

  @Override
  public E peek() {
    if (logicalSize == 0) {
      throw new RuntimeException("[PANIC - NoSuchElement]");
    }

    return data[logicalSize - 1];
  }

  @Override
  @SuppressWarnings("unchecked")
  public void push(E element) {
    if (logicalSize == physicalSize) {
      int newCapacity = DEFAULT_CAPACITY;

      if (physicalSize >= DEFAULT_CAPACITY) {
        newCapacity = physicalSize + (physicalSize >>> 1);
      }

      E[] temp = (E[]) new Object[newCapacity];

      for (int i = 0, length = logicalSize; i < length; i++) {
        temp[i] = data[i];
      }

      data = temp;
      physicalSize = newCapacity;
    }

    data[logicalSize] = element;

    logicalSize++;
  }

  @Override
  public void pop() {
    if (logicalSize == 0) {
      throw new RuntimeException("[PANIC - NoSuchElement]");
    }

    logicalSize--;

    data[logicalSize] = null;
  }

  public int capacity() {
    return physicalSize;
  }

  @SuppressWarnings("unchecked")
  public void shrink() {
    E[] temp = (E[]) new Object[logicalSize];

    for (int i = 0, length = logicalSize; i < length; i++) {
      temp[i] = data[i];
    }

    data = temp;
    physicalSize = logicalSize;
  }
}
