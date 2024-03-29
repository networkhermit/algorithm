package muse.data_structures;

import muse.data_structures.List;

public class CircularlyLinkedList<E> implements List<E> {
  private static class Node<E> {
    public E data;
    public Node<E> next = null;

    public Node(E element) {
      data = element;
    }
  }

  private Node<E> tail = null;
  private int length = 0;

  @Override
  public int size() {
    return length;
  }

  @Override
  public boolean isEmpty() {
    return length == 0;
  }

  @Override
  public E get(int index) {
    if (index < 0 || index >= length) {
      throw new RuntimeException("[PANIC - IndexOutOfBounds]");
    }

    Node<E> cursor = tail;

    if (index != length - 1) {
      for (int i = 0; i <= index; i++) {
        cursor = cursor.next;
      }
    }

    return cursor.data;
  }

  @Override
  public void set(int index, E element) {
    if (index < 0 || index >= length) {
      throw new RuntimeException("[PANIC - IndexOutOfBounds]");
    }

    Node<E> cursor = tail;

    if (index != length - 1) {
      for (int i = 0; i <= index; i++) {
        cursor = cursor.next;
      }
    }

    cursor.data = element;
  }

  @Override
  @SuppressWarnings("unchecked")
  public void insert(int index, E element) {
    if (index < 0 || index > length) {
      throw new RuntimeException("[PANIC - IndexOutOfBounds]");
    }

    Node<E> node = new Node<>(element);

    if (index == 0) {
      if (length == 0) {
        node.next = node;
        tail = node;
      } else {
        node.next = tail.next;
        tail.next = node;
      }
    } else if (index == length) {
      node.next = tail.next;
      tail.next = node;
      tail = node;
    } else {
      Node<E> cursor = tail;
      for (int i = 0, bound = index - 1; i <= bound; i++) {
        cursor = cursor.next;
      }
      node.next = cursor.next;
      cursor.next = node;
    }

    length++;
  }

  @Override
  public void remove(int index) {
    if (index < 0 || index >= length) {
      throw new RuntimeException("[PANIC - IndexOutOfBounds]");
    }

    Node<E> target;

    if (index == 0) {
      target = tail.next;
      if (length == 1) {
        tail = null;
      } else {
        tail.next = target.next;
      }
    } else {
      Node<E> cursor = tail;
      for (int i = 0, bound = index - 1; i <= bound; i++) {
        cursor = cursor.next;
      }
      target = cursor.next;
      cursor.next = target.next;
      if (index == length - 1) {
        tail = cursor;
      }
    }

    target.data = null;

    length--;
  }

  @Override
  public E front() {
    return get(0);
  }

  @Override
  public E back() {
    return get(length - 1);
  }

  @Override
  public void prepend(E element) {
    insert(0, element);
  }

  @Override
  public void append(E element) {
    insert(length, element);
  }

  @Override
  public void poll() {
    remove(0);
  }

  @Override
  public void eject() {
    remove(length - 1);
  }
}
