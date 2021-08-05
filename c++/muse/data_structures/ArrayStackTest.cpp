#include <muse/data_structures/ArrayStack.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testArrayStack() {
    size_t size = 8192;

    auto stack = new ArrayStack::ArrayStack<int>(0);

    for (size_t i = 1; i <= size; i++) {
        stack->push(static_cast<int>(i));
    }

    stack->shrink();

    if (stack->size() != size) {
        return false;
    }

    if (stack->capacity() != size) {
        return false;
    }

    for (size_t i = size; i > 0; i--) {
        if (static_cast<size_t>(stack->peek()) != i) {
            return false;
        }
        stack->pop();
    }

    stack->shrink();

    if (!stack->isEmpty()) {
        return false;
    }

    return stack->capacity() == 0;
}

int main() {
    TestRunner::pick(&testArrayStack);
}
