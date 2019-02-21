#include <fun/vac/data_structures/CircularlyLinkedList.h>
#include <fun/vac/util/TestRunner.h>

bool testCircularlyLinkedList(void) {
    size_t size = 8192;

    CircularlyLinkedList *list = CircularlyLinkedList_new();

    for (size_t i = 1; i <= size; i++) {
        CircularlyLinkedList_append(list, (int) i);
    }

    if (CircularlyLinkedList_size(list) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) CircularlyLinkedList_get(list, i) != i + 1) {
            return false;
        }
    }

    for (size_t i = 0; i < size; i++) {
        CircularlyLinkedList_set(list, i, size - i);
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) CircularlyLinkedList_get(list, i) != size - i) {
            return false;
        }
    }

    int x, y;

    for (size_t i = 0, j = size - 1; i < j; i++, j--) {
        x = CircularlyLinkedList_get(list, i);
        y = CircularlyLinkedList_get(list, j);

        CircularlyLinkedList_remove(list, i);
        CircularlyLinkedList_insert(list, i, y);
        CircularlyLinkedList_remove(list, j);
        CircularlyLinkedList_insert(list, j, x);
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) CircularlyLinkedList_get(list, i) != i + 1) {
            return false;
        }
    }

    for (size_t i = size; i >= 1; i--) {
        if ((size_t) CircularlyLinkedList_back(list) != i) {
            return false;
        } else {
            CircularlyLinkedList_eject(list);
        }
    }

    if (!CircularlyLinkedList_isEmpty(list)) {
        return false;
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testCircularlyLinkedList());
}
