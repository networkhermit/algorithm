package muse.data_structures;

import muse.data_structures.interfaces.IList;

public class LinkedList<E> implements IList<E> {
    private static class Node<E> {
        public E data;
        public Node<E> next = null;

        public Node(E element) {
            data = element;
        }
    }

    private Node<E> head = null;
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

        Node<E> cursor;

        if (index == length - 1) {
            cursor = tail;
        } else {
            cursor = head;
            for (int i = 0; i < index; i++) {
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

        Node<E> cursor;

        if (index == length - 1) {
            cursor = tail;
        } else {
            cursor = head;
            for (int i = 0; i < index; i++) {
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
            if (length != 0) {
                node.next = head;
            } else {
                tail = node;
            }
            head = node;
        } else if (index == length) {
            tail.next = node;
            tail = node;
        } else {
            Node<E> cursor = head;
            for (int i = 0, bound = index - 1; i < bound; i++) {
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
            target = head;
            if (length == 1) {
                head = null;
                tail = null;
            } else {
                head = target.next;
            }
        } else {
            Node<E> cursor = head;
            for (int i = 0, bound = index - 1; i < bound; i++) {
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
