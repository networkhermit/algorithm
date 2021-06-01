#include <muse/data_structures/ArrayQueue.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testArrayQueue() {
    size_t size = 8192;

    auto queue = new ArrayQueue::ArrayQueue<int>(0);

    for (size_t i = 1; i <= size; i++) {
        queue->enqueue(static_cast<int>(i));
    }

    queue->shrink();

    if (queue->size() != size) {
        return false;
    }

    if (queue->capacity() != size) {
        return false;
    }

    for (size_t i = 1; i <= size; i++) {
        if (static_cast<size_t>(queue->peek()) != i) {
            return false;
        }
        queue->dequeue();
    }

    queue->shrink();

    if (!queue->isEmpty()) {
        return false;
    }

    return queue->capacity() == 0;
}

int main() {
    TestRunner::parseTest(testArrayQueue());
}
