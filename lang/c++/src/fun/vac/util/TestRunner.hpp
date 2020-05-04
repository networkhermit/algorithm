#ifndef FUN_VAC_UTIL_TEST_RUNNER_HPP
#define FUN_VAC_UTIL_TEST_RUNNER_HPP 1

#include <cstddef>
#include <iostream>

namespace TestRunner {

    void parseTest(bool ok) {
        static std::size_t TestRunnerItemIndex = 0;

        if (ok) {
            std::cout << "✓ Item [" << TestRunnerItemIndex << "] PASSED" << std::endl;
        } else {
            std::cerr << "✗ Item [" << TestRunnerItemIndex << "] FAILED" << std::endl;
        }

        TestRunnerItemIndex++;
    }
}

#endif
