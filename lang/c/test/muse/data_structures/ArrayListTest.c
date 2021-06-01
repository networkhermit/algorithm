#include <muse/data_structures/ArrayList.h>
#include <muse/util/TestRunner.h>

bool testArrayList(void) {
    size_t size = 8192;

    ArrayList *list = ArrayList_new(0);

    for (size_t i = 1; i <= size; i++) {
        ArrayList_append(list, (int) i);
    }

    ArrayList_shrink(list);

    if (ArrayList_size(list) != size) {
        return false;
    }

    if (ArrayList_capacity(list) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) ArrayList_get(list, i) != i + 1) {
            return false;
        }
    }

    for (size_t i = 0; i < size; i++) {
        ArrayList_set(list, i, size - i);
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) ArrayList_get(list, i) != size - i) {
            return false;
        }
    }

    int x;
    int y;

    for (size_t i = 0, j = size - 1; i < j; i++, j--) {
        x = ArrayList_get(list, i);
        y = ArrayList_get(list, j);

        ArrayList_remove(list, i);
        ArrayList_insert(list, i, y);
        ArrayList_remove(list, j);
        ArrayList_insert(list, j, x);
    }

    for (size_t i = 0; i < size; i++) {
        if ((size_t) ArrayList_get(list, i) != i + 1) {
            return false;
        }
    }

    for (size_t i = size; i >= 1; i--) {
        if ((size_t) ArrayList_back(list) != i) {
            return false;
        }
        ArrayList_eject(list);
    }

    ArrayList_shrink(list);

    if (!ArrayList_isEmpty(list)) {
        return false;
    }

    return ArrayList_capacity(list) == 0;
}

int main(void) {
    TestRunner_parseTest(testArrayList());
}
