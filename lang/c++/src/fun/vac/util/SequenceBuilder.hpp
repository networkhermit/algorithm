#ifndef FUN_VAC_UTIL_SEQUENCE_BUILDER_HPP
#define FUN_VAC_UTIL_SEQUENCE_BUILDER_HPP 1

#include <cstddef>
#include <vector>

#include <fun/vac/util/RandomFactory.hpp>
#include <fun/vac/util/Sequences.hpp>

namespace SequenceBuilder {

    void packIncreasing(std::vector<int> &arr) {
        RandomFactory::seed();
        arr[0] = RandomFactory::integerN(1, 4);
        for (std::size_t i = 1, length = arr.size(); i < length; i++) {
            arr[i] = arr[i - 1] + RandomFactory::integerN(1, 4);
        }
    }

    void packRandom(std::vector<int> &arr) {
        RandomFactory::seed();
        for (std::size_t i = 0, length = arr.size(); i < length; i++) {
            arr[i] = RandomFactory::generateInteger();
        }
    }

    void packDecreasing(std::vector<int> &arr) {
        packIncreasing(arr);
        Sequences::reverse(arr);
    }
}

#endif
