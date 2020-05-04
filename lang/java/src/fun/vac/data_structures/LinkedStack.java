package fun.vac.data_structures;

import fun.vac.data_structures.interfaces.IStack;

public class LinkedStack<E> implements IStack<E> {

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

    public LinkedStack() {}

    @Override
    public int size() {
        return length;
    }

    @Override
    public boolean isEmpty() {
        return length == 0;
    }

    @Override
    public E peek() {
        if (length == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        return tail.data;
    }

    @Override
    @SuppressWarnings("unchecked")
    public void push(E element) {
        Node<E> node = new Node<>(element);

        if (length == 0) {
            head = node;
        } else {
            tail.next = node;
        }

        tail = node;

        length++;
    }

    @Override
    public void pop() {
        if (length == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        Node<E> target = tail;

        if (length == 1) {
            head = null;
            tail = null;
        } else {
            Node<E> cursor = head;
            for (int i = 0, bound = length - 2; i < bound; i++) {
                cursor = cursor.next;
            }
            cursor.next = null;
            tail = cursor;
        }

        target.data = null;

        length--;
    }
}
