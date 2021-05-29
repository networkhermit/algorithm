#include <vac.fun/data_structures/LinkedStack.hpp>
#include <vac.fun/util/TestRunner.hpp>

using namespace std;

bool testLinkedStack() {
    size_t size = 8192;

    auto stack = new LinkedStack::LinkedStack<int>();

    for (size_t i = 1; i <= size; i++) {
        stack->push(static_cast<int>(i));
    }

    if (stack->size() != size) {
        return false;
    }

    for (size_t i = size; i > 0; i--) {
        if (static_cast<size_t>(stack->peek()) != i) {
            return false;
        }
        stack->pop();
    }

    return stack->isEmpty();
}

int main() {
    TestRunner::parseTest(testLinkedStack());
}
