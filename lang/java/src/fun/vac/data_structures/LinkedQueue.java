package fun.vac.data_structures;

import fun.vac.data_structures.interfaces.IQueue;

public class LinkedQueue<E> implements IQueue<E> {

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

    public LinkedQueue() {}

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

        return head.data;
    }

    @Override
    @SuppressWarnings("unchecked")
    public void enqueue(E element) {
        Node<E> node = new Node<>(element);

        if (length == 0) {
            head = node;
        } else {
            tail.next = node;
        }

        tail = node;

        length += 1;
    }

    @Override
    public void dequeue() {
        if (length == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        Node<E> target = head;

        if (length == 1) {
            head = null;
            tail = null;
        } else {
            head = target.next;
        }

        target.data = null;

        length -= 1;
    }
}
