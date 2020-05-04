package fun.vac.data_structures;

import fun.vac.data_structures.interfaces.IList;

public class DoublyLinkedList<E> implements IList<E> {

    private static class Node<E> {

        public E data;
        public Node<E> next = null;
        public Node<E> prev = null;

        public Node(E element) {
            data = element;
        }
    }

    private Node<E> head = null;
    private Node<E> tail = null;
    private int length = 0;

    public DoublyLinkedList() {}

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

        Node<E> cursor;

        if (index < length >>> 1) {
            cursor = head;
            for (int i = 0; i < index; i++) {
                cursor = cursor.next;
            }
        } else {
            cursor = tail;
            for (int i = length - 1; i > index; i--) {
                cursor = cursor.prev;
            }
        }

        return cursor.data;
    }

    @Override
    public void set(int index, E element) {
        if (index < 0 || index >= length) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        Node<E> cursor;

        if (index < length >>> 1) {
            cursor = head;
            for (int i = 0; i < index; i++) {
                cursor = cursor.next;
            }
        } else {
            cursor = tail;
            for (int i = length - 1; i > index; i--) {
                cursor = cursor.prev;
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
            if (length != 0) {
                node.next = head;
                head.prev = node;
            } else {
                tail = node;
            }
            head = node;
        } else if (index == length) {
            node.prev = tail;
            tail.next = node;
            tail = node;
        } else {
            Node<E> cursor;
            if (index < length >>> 1) {
                cursor = head;
                for (int i = 0; i < index; i++) {
                    cursor = cursor.next;
                }
            } else {
                cursor = tail;
                for (int i = length - 1; i > index; i--) {
                    cursor = cursor.prev;
                }
            }
            node.next = cursor;
            node.prev = cursor.prev;
            cursor.prev.next = node;
            cursor.prev = node;
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
            target = head;
            if (length == 1) {
                head = null;
                tail = null;
            } else {
                target.next.prev = null;
                head = target.next;
            }
        } else if (index == length - 1) {
            target = tail;
            target.prev.next = null;
            tail = target.prev;
        } else {
            if (index < length >>> 1) {
                target = head;
                for (int i = 0; i < index; i++) {
                    target = target.next;
                }
            } else {
                target = tail;
                for (int i = length - 1; i > index; i--) {
                    target = target.prev;
                }
            }
            target.prev.next = target.next;
            target.next.prev = target.prev;
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
