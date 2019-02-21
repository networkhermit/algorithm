package fun.vac.data_structures.interfaces;

public interface IStack<E> {

    int size();

    boolean isEmpty();

    E peek();

    void push(E element);

    void pop();
}
