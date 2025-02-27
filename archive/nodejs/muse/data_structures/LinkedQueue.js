class Node {
  constructor(element) {
    this.data = element;
    this.next = null;
  }
}

export class LinkedQueue {
  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  size() {
    return this.length;
  }

  isEmpty() {
    return this.length === 0;
  }

  peek() {
    if (this.length === 0) {
      throw new Error('[PANIC - NoSuchElement]');
    }

    return this.head.data;
  }

  enqueue(element) {
    const node = new Node(element);

    if (this.length === 0) {
      this.head = node;
    } else {
      this.tail.next = node;
    }

    this.tail = node;

    this.length++;
  }

  dequeue() {
    if (this.length === 0) {
      throw new Error('[PANIC - NoSuchElement]');
    }

    const target = this.head;

    if (this.length === 1) {
      this.head = null;
      this.tail = null;
    } else {
      this.head = target.next;
    }

    target.data = null;

    this.length--;
  }
}
