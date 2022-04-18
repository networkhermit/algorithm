#ifndef MUSE_DATA_STRUCTURES_LINKED_QUEUE_HPP
#define MUSE_DATA_STRUCTURES_LINKED_QUEUE_HPP

#include <stdexcept>

namespace LinkedQueue {
template <typename E> class LinkedQueue {
private:
  class Node {
  public:
    E data;
    Node *next = nullptr;

  public:
    Node(E element) { data = element; }

    ~Node() {}
  };

private:
  Node *head = nullptr;
  Node *tail = nullptr;
  std::size_t length = 0;

public:
  LinkedQueue() {}

  ~LinkedQueue() {}

public:
  std::size_t size() { return length; }

  bool isEmpty() { return length == 0; }

  E peek() {
    if (length == 0) {
      throw std::runtime_error("[PANIC - NoSuchElement]");
    }

    return head->data;
  }

  void enqueue(E element) {
    auto node = new Node(element);

    if (length == 0) {
      head = node;
    } else {
      tail->next = node;
    }

    tail = node;

    length++;
  }

  void dequeue() {
    if (length == 0) {
      throw std::runtime_error("[PANIC - NoSuchElement]");
    }

    auto target = head;

    if (length == 1) {
      head = nullptr;
      tail = nullptr;
    } else {
      head = target->next;
    }

    target->data = static_cast<E>(0);
    delete target;

    length--;
  }
};
} // namespace LinkedQueue

#endif
