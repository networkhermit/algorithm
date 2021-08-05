#ifndef MUSE_UTIL_TEST_RUNNER_HPP
#define MUSE_UTIL_TEST_RUNNER_HPP 1

#include <cstddef>
#include <functional>
#include <iostream>

namespace TestRunner {

    void pick(std::function<bool()> func) {
        static std::size_t TestRunnerItemIndex = 0;

        if (func()) {
            std::cout << "✓ Item [" << TestRunnerItemIndex << "] PASSED" << std::endl;
        } else {
            std::cerr << "✗ Item [" << TestRunnerItemIndex << "] FAILED" << std::endl;
        }

        TestRunnerItemIndex++;
    }
}

#endif
