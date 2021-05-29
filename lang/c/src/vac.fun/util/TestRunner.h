#ifndef VAC_FUN_UTIL_TEST_RUNNER_H
#define VAC_FUN_UTIL_TEST_RUNNER_H 1

#include <stdbool.h>
#include <stdio.h>

void TestRunner_parseTest(bool ok) {
    static size_t TestRunnerItemIndex = 0;

    if (ok) {
        printf("✓ Item [%zu] PASSED\n", TestRunnerItemIndex);
    } else {
        fprintf(stderr, "✗ Item [%zu] FAILED\n", TestRunnerItemIndex);
    }

    TestRunnerItemIndex++;
}

#endif
