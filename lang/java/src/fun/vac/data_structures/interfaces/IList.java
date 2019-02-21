package fun.vac.data_structures.interfaces;

public interface IList<E> {

    int size();

    boolean isEmpty();

    E get(int index);

    void set(int index, E element);

    void insert(int index, E element);

    void remove(int index);

    E front();

    E back();

    void prepend(E element);

    void append(E element);

    void poll();

    void eject();
}
