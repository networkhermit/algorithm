#include <muse/data_structures/DoublyLinkedList.h>
#include <muse/util/TestRunner.h>

bool testLinkedList(void) {
    size_t size = 8192;

    DoublyLinkedList *list = DoublyLinkedList_new();

    for (size_t i = 1; i <= size; i++) {
        DoublyLinkedList_append(list, (int) i);
    }

    if (DoublyLinkedList_size(list) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) DoublyLinkedList_get(list, i) != i + 1) {
            return false;
        }
    }

    for (size_t i = 0; i < size; i++) {
        DoublyLinkedList_set(list, i, size - i);
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) DoublyLinkedList_get(list, i) != size - i) {
            return false;
        }
    }

    int x;
    int y;

    for (size_t i = 0, j = size - 1; i < j; i++, j--) {
        x = DoublyLinkedList_get(list, i);
        y = DoublyLinkedList_get(list, j);

        DoublyLinkedList_remove(list, i);
        DoublyLinkedList_insert(list, i, y);
        DoublyLinkedList_remove(list, j);
        DoublyLinkedList_insert(list, j, x);
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) DoublyLinkedList_get(list, i) != i + 1) {
            return false;
        }
    }

    for (size_t i = size; i >= 1; i--) {
        if ((size_t) DoublyLinkedList_back(list) != i) {
            return false;
        }
        DoublyLinkedList_eject(list);
    }

    return DoublyLinkedList_isEmpty(list);
}

int main(void) {
    TestRunner_pick(&testLinkedList);
}
