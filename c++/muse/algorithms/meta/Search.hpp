#ifndef MUSE_ALGORITHMS_META_SEARCH_HPP
#define MUSE_ALGORITHMS_META_SEARCH_HPP

#include <muse/algorithms/search/BinarySearch.hpp>
#include <muse/algorithms/search/LinearSearch.hpp>

namespace Search {

    template <typename T>
    std::size_t binarySearch(const std::vector<T> &arr, T key) {
        return BinarySearch::find(arr, key);
    }

    template <typename T>
    std::size_t linearSearch(const std::vector<T> &arr, T key) {
        return LinearSearch::find(arr, key);
    }
}

#endif
