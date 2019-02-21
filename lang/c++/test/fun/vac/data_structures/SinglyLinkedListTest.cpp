#include <fun/vac/data_structures/SinglyLinkedList.hpp>
#include <fun/vac/util/TestRunner.hpp>

using namespace std;

bool testSinglyLinkedList() {
    size_t size = 8192;

    auto list = new SinglyLinkedList::SinglyLinkedList<int>();

    for (size_t i = 1; i <= size; i++) {
        list->append(static_cast<int>(i));
    }

    if (list->size() != size) {
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

    int x, y;

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
        } else {
            list->eject();
        }
    }

    if (!list->isEmpty()) {
        return false;
    }

    return true;
}

int main() {
    TestRunner::parseTest(testSinglyLinkedList());
}
