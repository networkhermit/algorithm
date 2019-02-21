#include <fun/vac/data_structures/LinkedQueue.h>
#include <fun/vac/util/TestRunner.h>

bool testLinkedQueue(void) {
    size_t size = 8192;

    LinkedQueue *queue = LinkedQueue_new();

    for (size_t i = 1; i <= size; i++) {
        LinkedQueue_enqueue(queue, (int) i);
    }

    if (LinkedQueue_size(queue) != size) {
        return false;
    }

    for (size_t i = 1; i <= size; i++) {
        if ((size_t) LinkedQueue_peek(queue) != i) {
            return false;
        }
        LinkedQueue_dequeue(queue);
    }

    if (!LinkedQueue_isEmpty(queue)) {
        return false;
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testLinkedQueue());
}
