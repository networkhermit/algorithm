#include <fun/vac/data_structures/ArrayQueue.h>
#include <fun/vac/util/TestRunner.h>

bool testArrayQueue(void) {
    size_t size = 8192;

    ArrayQueue *queue = ArrayQueue_new(0);

    for (size_t i = 1; i <= size; i++) {
        ArrayQueue_enqueue(queue, (int) i);
    }

    ArrayQueue_shrink(queue);

    if (ArrayQueue_size(queue) != size) {
        return false;
    }

    if (ArrayQueue_capacity(queue) != size) {
        return false;
    }

    for (size_t i = 1; i <= size; i++) {
        if ((size_t) ArrayQueue_peek(queue) != i) {
            return false;
        }
        ArrayQueue_dequeue(queue);
    }

    ArrayQueue_shrink(queue);

    if (!ArrayQueue_isEmpty(queue)) {
        return false;
    }

    if (ArrayQueue_capacity(queue) != 0) {
        return false;
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testArrayQueue());
}
