#ifndef FUN_VAC_UTIL_TEST_RUNNER_HPP
#define FUN_VAC_UTIL_TEST_RUNNER_HPP 1

#include <cstddef>
#include <iostream>

namespace TestRunner {

    void parseTest(bool ok) {
        static std::size_t TestRunner_TestIndex = 0;

        if (ok) {
            std::cout << "< Test [" << TestRunner_TestIndex << "] Passed >" << std::endl;
        } else {
            std::cerr << "X Test [" << TestRunner_TestIndex << "] Failed X" << std::endl;
        }

        TestRunner_TestIndex += 1;
    }
}

#endif
