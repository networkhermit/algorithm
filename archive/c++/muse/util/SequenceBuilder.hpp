#ifndef MUSE_UTIL_SEQUENCE_BUILDER_HPP
#define MUSE_UTIL_SEQUENCE_BUILDER_HPP

#include <cstddef>
#include <vector>

#include <muse/util/RandomFactory.hpp>
#include <muse/util/Sequences.hpp>

namespace SequenceBuilder {

void packIncreasing(std::vector<int> &arr) {
  RandomFactory::seed();
  arr[0] = RandomFactory::genIntN(1, 3);
  for (std::size_t i = 1, length = arr.size(); i < length; i++) {
    arr[i] = arr[i - 1] + RandomFactory::genIntN(1, 3);
  }
}

void packRandom(std::vector<int> &arr) {
  RandomFactory::seed();
  for (std::size_t i = 0, length = arr.size(); i < length; i++) {
    arr[i] = RandomFactory::genInt();
  }
}

void packDecreasing(std::vector<int> &arr) {
  packIncreasing(arr);
  Sequences::reverse(arr);
}
} // namespace SequenceBuilder

#endif
