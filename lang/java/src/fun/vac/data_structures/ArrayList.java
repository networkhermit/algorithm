package fun.vac.data_structures;

import fun.vac.data_structures.interfaces.IList;

public class ArrayList<E> implements IList<E> {

    private static final int DEFAULT_CAPACITY = 64;

    private E[] data = null;
    private int logicalSize = 0;
    private int physicalSize = DEFAULT_CAPACITY;

    public ArrayList() {
        this(0);
    }

    @SuppressWarnings("unchecked")
    public ArrayList(int physicalSize) {
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
    public E get(int index) {
        if (index < 0 || index >= logicalSize) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        return data[index];
    }

    @Override
    public void set(int index, E element) {
        if (index < 0 || index >= logicalSize) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        data[index] = element;
    }

    @Override
    @SuppressWarnings("unchecked")
    public void insert(int index, E element) {
        if (index < 0 || index > logicalSize) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        if (logicalSize == physicalSize) {
            int newCapacity = DEFAULT_CAPACITY;

            if (physicalSize >= DEFAULT_CAPACITY) {
                newCapacity = physicalSize + (physicalSize >>> 1);
            }

            E[] temp = (E[]) new Object[newCapacity];

            for (int i = 0, length = logicalSize; i < length; i++) {
                temp[i] = data[i];
            }

            data = temp;
            physicalSize = newCapacity;
        }

        for (int i = logicalSize; i > index; i--) {
            data[i] = data[i - 1];
        }

        data[index] = element;

        logicalSize++;
    }

    @Override
    public void remove(int index) {
        if (index < 0 || index >= logicalSize) {
            throw new RuntimeException("[PANIC - IndexOutOfBounds]");
        }

        for (int i = index + 1; i < logicalSize; i++) {
            data[i - 1] = data[i];
        }

        logicalSize--;

        data[logicalSize] = null;
    }

    @Override
    public E front() {
        return get(0);
    }

    @Override
    public E back() {
        return get(logicalSize - 1);
    }

    @Override
    public void prepend(E element) {
        insert(0, element);
    }

    @Override
    public void append(E element) {
        insert(logicalSize, element);
    }

    @Override
    public void poll() {
        remove(0);
    }

    @Override
    public void eject() {
        remove(logicalSize - 1);
    }

    public int capacity() {
        return physicalSize;
    }

    @SuppressWarnings("unchecked")
    public void shrink() {
        E[] temp = (E[]) new Object[logicalSize];

        for (int i = 0, length = logicalSize; i < length; i++) {
            temp[i] = data[i];
        }

        data = temp;
        physicalSize = logicalSize;
    }
}
