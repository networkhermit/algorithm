package fun.vac.data_structures;

import fun.vac.data_structures.interfaces.IQueue;

public class ArrayQueue<E> implements IQueue<E> {

    private static final int DEFAULT_CAPACITY = 64;

    private E[] data = null;
    private int front = 0;
    private int logicalSize = 0;
    private int physicalSize = DEFAULT_CAPACITY;

    public ArrayQueue() {
        this(0);
    }

    @SuppressWarnings("unchecked")
    public ArrayQueue(int physicalSize) {
        if (physicalSize > 1) {
            this.physicalSize = physicalSize;
        }
        data = (E[]) new Object[this.physicalSize];
    }

    @Override
    public int size() {
        return logicalSize;
    }

    @Override
    public boolean isEmpty() {
        return logicalSize == 0;
    }

    @Override
    public E peek() {
        if (logicalSize == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        return data[front];
    }

    @Override
    @SuppressWarnings("unchecked")
    public void enqueue(E element) {
        if (logicalSize == physicalSize) {
            int newCapacity = DEFAULT_CAPACITY;

            if (physicalSize >= DEFAULT_CAPACITY) {
                newCapacity = physicalSize + (physicalSize >>> 1);
            }

            E[] temp = (E[]) new Object[newCapacity];

            int cursor = front;

            for (int i = 0, length = logicalSize; i < length; i++) {
                if (cursor == physicalSize) {
                    cursor = 0;
                }
                temp[i] = data[cursor];
                cursor++;
            }

            data = temp;
            front = 0;
            physicalSize = newCapacity;
        }

        data[(front + logicalSize) % physicalSize] = element;

        logicalSize++;
    }

    @Override
    public void dequeue() {
        if (logicalSize == 0) {
            throw new RuntimeException("[PANIC - NoSuchElement]");
        }

        data[front] = null;

        front = (front + 1) % physicalSize;

        logicalSize--;
    }

    public int capacity() {
        return physicalSize;
    }

    @SuppressWarnings("unchecked")
    public void shrink() {
        E[] temp = (E[]) new Object[logicalSize];

        int cursor = front;

        for (int i = 0, length = logicalSize; i < length; i++) {
            if (cursor == physicalSize) {
                cursor = 0;
            }
            temp[i] = data[cursor];
            cursor++;
        }

        data = temp;
        front = 0;
        physicalSize = logicalSize;
    }
}
