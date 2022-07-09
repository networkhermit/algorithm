package muse.data_structures;

public interface Queue<E> {
  int size();

  boolean isEmpty();

  E peek();

  void enqueue(E element);

  void dequeue();
}
