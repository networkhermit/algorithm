#ifndef FUN_VAC_UTIL_TEST_RUNNER_H
#define FUN_VAC_UTIL_TEST_RUNNER_H 1

#include <stdbool.h>
#include <stdio.h>

void TestRunner_parseTest(bool ok) {
    static size_t TestRunnerItemIndex = 0;

    if (ok) {
        printf("✓ Item [%d] PASSED\n", TestRunnerItemIndex);
    } else {
        fprintf(stderr, "✗ Item [%d] FAILED\n", TestRunnerItemIndex);
    }

    TestRunnerItemIndex++;
}

#endif
