#include <muse/data_structures/ArrayList.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testArrayList() {
    size_t size = 8192;

    auto list = new ArrayList::ArrayList<int>(0);

    for (size_t i = 1; i <= size; i++) {
        list->append(static_cast<int>(i));
    }

    list->shrink();

    if (list->size() != size) {
        return false;
    }

    if (list->capacity() != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (static_cast<size_t>(list->get(i)) != i + 1) {
            return false;
        }
    }

    for (size_t i = 0; i < size; i++) {
        list->set(i, size - i);
    }

    for (size_t i = 0; i < size; i++) {
        if (static_cast<size_t>(list->get(i)) != size - i) {
            return false;
        }
    }

    int x;
    int y;

    for (size_t i = 0, j = size - 1; i < j; i++, j--) {
        x = list->get(i);
        y = list->get(j);

        list->remove(i);
        list->insert(i, y);
        list->remove(j);
        list->insert(j, x);
    }

    for (size_t i = 0; i < size; i++) {
        if (static_cast<size_t>(list->get(i)) != i + 1) {
            return false;
        }
    }

    for (size_t i = size; i >= 1; i--) {
        if (static_cast<size_t>(list->back()) != i) {
            return false;
        }
        list->eject();
    }

    list->shrink();

    if (!list->isEmpty()) {
        return false;
    }

    return list->capacity() == 0;
}

int main() {
    TestRunner::pick(&testArrayList);
}
