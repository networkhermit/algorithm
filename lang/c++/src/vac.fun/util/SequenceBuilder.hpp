#ifndef VAC_FUN_UTIL_SEQUENCE_BUILDER_HPP
#define VAC_FUN_UTIL_SEQUENCE_BUILDER_HPP 1

#include <cstddef>
#include <vector>

#include <vac.fun/util/RandomFactory.hpp>
#include <vac.fun/util/Sequences.hpp>

namespace SequenceBuilder {

    void packIncreasing(std::vector<int> &arr) {
        RandomFactory::seed();
        arr[0] = RandomFactory::integerN(1, 3);
        for (std::size_t i = 1, length = arr.size(); i < length; i++) {
            arr[i] = arr[i - 1] + RandomFactory::integerN(1, 3);
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
