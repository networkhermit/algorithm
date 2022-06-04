package muse.data_structures;

public interface Stack<E> {
  int size();

  boolean isEmpty();

  E peek();

  void push(E element);

  void pop();
}
