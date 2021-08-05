#ifndef MUSE_UTIL_TEST_RUNNER_H
#define MUSE_UTIL_TEST_RUNNER_H 1

#include <stdbool.h>
#include <stdio.h>

void TestRunner_pick(bool (*func)()) {
    static size_t TestRunnerItemIndex = 0;

    if (func()) {
        printf("✓ Item [%zu] PASSED\n", TestRunnerItemIndex);
    } else {
        fprintf(stderr, "✗ Item [%zu] FAILED\n", TestRunnerItemIndex);
    }

    TestRunnerItemIndex++;
}

#endif
