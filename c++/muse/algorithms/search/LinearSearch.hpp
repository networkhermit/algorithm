#ifndef MUSE_ALGORITHMS_SEARCH_LINEAR_SEARCH_HPP
#define MUSE_ALGORITHMS_SEARCH_LINEAR_SEARCH_HPP

#include <cstddef>
#include <vector>

namespace LinearSearch {

    template <typename T>
    std::size_t find(const std::vector<T> &arr, T key) {
        for (std::size_t i = 0, length = arr.size(); i < length; i++) {
            if (arr[i] == key) {
                return i;
            }
        }

        return arr.size();
    }
}

#endif
