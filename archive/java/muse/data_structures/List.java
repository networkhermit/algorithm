package muse.data_structures;

public interface List<E> {
  int size();

  boolean isEmpty();

  E get(int index);

  void set(int index, E element);

  void insert(int index, E element);

  void remove(int index);

  E front();

  E back();

  void prepend(E element);

  void append(E element);

  void poll();

  void eject();
}
