#ifndef FUN_VAC_UTIL_TEST_RUNNER_H
#define FUN_VAC_UTIL_TEST_RUNNER_H 1

#include <stdbool.h>
#include <stdio.h>

void TestRunner_parseTest(bool ok) {
    static size_t TestRunner_TestIndex = 0;

    if (ok) {
        printf("✓ Test [%d] Passed\n", TestRunner_TestIndex);
    } else {
        fprintf(stderr, "✗ Test [%d] Failed\n", TestRunner_TestIndex);
    }

    TestRunner_TestIndex += 1;
}

#endif
