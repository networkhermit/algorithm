#ifndef VAC_FUN_UTIL_SEQUENCES_HPP
#define VAC_FUN_UTIL_SEQUENCES_HPP 1

#include <cstddef>
#include <iomanip>
#include <iostream>
#include <vector>

#include <vac.fun/algorithms/sorting/QuickSort.hpp>
#include <vac.fun/util/RandomFactory.hpp>

namespace Sequences {

    void inspect(const std::vector<int> &arr) {
        std::cout << "[" << std::endl;
        for (std::size_t i = 0, length = arr.size(); i < length; i++) {
            std::cout << "\t#" << std::setfill('0') << std::setw(4) << std::uppercase << std::hex << i << "  -->  " << std::dec << arr[i] << std::endl;
        }
        std::cout << "]" << std::endl;
    }

    bool isSorted(const std::vector<int> &arr) {
        for (std::size_t i = 1, length = arr.size(); i < length; i++) {
            if (arr[i] < arr[i - 1]) {
                return false;
            }
        }

        return true;
    }

    int parityChecksum(const std::vector<int> &arr) {
        int checksum = 0;

        for (std::size_t i = 0, length = arr.size(); i < length; i++) {
            checksum ^= arr[i];
        }

        return checksum;
    }

    void reverse(std::vector<int> &arr) {
        int k;
        int temp;

        for (std::size_t i = 0, bound = arr.size() >> 1, length = arr.size(); i < bound; i++) {
            k = length - i - 1;
            temp = arr[i];
            arr[i] = arr[k];
            arr[k] = temp;
        }
    }

    void shuffle(std::vector<int> &arr) {
        int k;
        int temp;

        RandomFactory::seed();
        for (std::size_t i = 0, length = arr.size(); i < length; i++) {
            k = RandomFactory::integerN(i, length);
            temp = arr[i];
            arr[i] = arr[k];
            arr[k] = temp;
        }
    }

    void sort(std::vector<int> &arr) {
        QuickSort::sort(arr);
    }
}

#endif
