package muse.data_structures.interfaces;

public interface IQueue<E> {
    int size();

    boolean isEmpty();

    E peek();

    void enqueue(E element);

    void dequeue();
}
