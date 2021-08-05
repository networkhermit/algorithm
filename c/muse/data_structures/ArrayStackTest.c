#include <muse/data_structures/ArrayStack.h>
#include <muse/util/TestRunner.h>

bool testArrayStack(void) {
    size_t size = 8192;

    ArrayStack *stack = ArrayStack_new(0);

    for (size_t i = 1; i <= size; i++) {
        ArrayStack_push(stack, (int) i);
    }

    ArrayStack_shrink(stack);

    if (ArrayStack_size(stack) != size) {
        return false;
    }

    if (ArrayStack_capacity(stack) != size) {
        return false;
    }

    for (size_t i = size; i > 0; i--) {
        if ((size_t) ArrayStack_peek(stack) != i) {
            return false;
        }
        ArrayStack_pop(stack);
    }

    ArrayStack_shrink(stack);

    if (!ArrayStack_isEmpty(stack)) {
        return false;
    }

    return ArrayStack_capacity(stack) == 0;
}

int main(void) {
    TestRunner_pick(&testArrayStack);
}
