#ifndef MUSE_UTIL_TEST_RUNNER_HPP
#define MUSE_UTIL_TEST_RUNNER_HPP

#include <cstddef>
#include <functional>
#include <iostream>

namespace TestRunner {

void pick(std::function<bool()> func) {
  static std::size_t itemIndex = 0;

  if (func()) {
    std::cout << "✓ Item [" << itemIndex << "] PASSED" << std::endl;
  } else {
    std::cerr << "✗ Item [" << itemIndex << "] FAILED" << std::endl;
  }

  itemIndex++;
}
} // namespace TestRunner

#endif
