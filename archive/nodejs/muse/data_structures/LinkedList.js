class Node {
  constructor(element) {
    this.data = element;
    this.next = null;
  }
}

export class LinkedList {
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

  get(index) {
    if (index < 0 || index >= this.length) {
      throw new Error('[PANIC - IndexOutOfBounds]');
    }

    let cursor = null;

    if (index === this.length - 1) {
      cursor = this.tail;
    } else {
      cursor = this.head;
      for (let i = 0; i < index; i++) {
        cursor = cursor.next;
      }
    }

    return cursor.data;
  }

  set(index, element) {
    if (index < 0 || index >= this.length) {
      throw new Error('[PANIC - IndexOutOfBounds]');
    }

    let cursor = null;

    if (index === this.length - 1) {
      cursor = this.tail;
    } else {
      cursor = this.head;
      for (let i = 0; i < index; i++) {
        cursor = cursor.next;
      }
    }

    cursor.data = element;
  }

  insert(index, element) {
    if (index < 0 || index > this.length) {
      throw new Error('[PANIC - IndexOutOfBounds]');
    }

    const node = new Node(element);

    switch (index) {
      case 0:
        if (this.length !== 0) {
          node.next = this.head;
        } else {
          this.tail = node;
        }
        this.head = node;
        break;
      case this.length:
        this.tail.next = node;
        this.tail = node;
        break;
      default: {
        let cursor = this.head;
        for (let i = 0, bound = index - 1; i < bound; i++) {
          cursor = cursor.next;
        }
        node.next = cursor.next;
        cursor.next = node;
      }
    }

    this.length++;
  }

  remove(index) {
    if (index < 0 || index >= this.length) {
      throw new Error('[PANIC - IndexOutOfBounds]');
    }

    let target = null;

    if (index === 0) {
      target = this.head;
      if (this.length === 1) {
        this.head = null;
        this.tail = null;
      } else {
        this.head = target.next;
      }
    } else {
      let cursor = this.head;
      for (let i = 0, bound = index - 1; i < bound; i++) {
        cursor = cursor.next;
      }
      target = cursor.next;
      cursor.next = target.next;
      if (index === this.length - 1) {
        this.tail = cursor;
      }
    }

    target.data = null;

    this.length--;
  }

  front() {
    return this.get(0);
  }

  back() {
    return this.get(this.length - 1);
  }

  prepend(element) {
    this.insert(0, element);
  }

  append(element) {
    this.insert(this.length, element);
  }

  poll() {
    this.remove(0);
  }

  eject() {
    this.remove(this.length - 1);
  }
}
