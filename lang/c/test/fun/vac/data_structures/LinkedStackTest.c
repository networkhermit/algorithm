#include <fun/vac/data_structures/LinkedStack.h>
#include <fun/vac/util/TestRunner.h>

bool testLinkedStack(void) {
    size_t size = 8192;

    LinkedStack *stack = LinkedStack_new();

    for (size_t i = 1; i <= size; i++) {
        LinkedStack_push(stack, (int) i);
    }

    if (LinkedStack_size(stack) != size) {
        return false;
    }

    for (size_t i = size; i > 0; i--) {
        if ((size_t) LinkedStack_peek(stack) != i) {
            return false;
        }
        LinkedStack_pop(stack);
    }

    return LinkedStack_isEmpty(stack);
}

int main(void) {
    TestRunner_parseTest(testLinkedStack());
}
