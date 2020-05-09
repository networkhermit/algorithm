#ifndef FUN_VAC_ALGORITHMS_META_SORTING_HPP
#define FUN_VAC_ALGORITHMS_META_SORTING_HPP 1

#include <fun/vac/algorithms/sorting/BubbleSort.hpp>
#include <fun/vac/algorithms/sorting/InsertionSort.hpp>
#include <fun/vac/algorithms/sorting/MergeSort.hpp>
#include <fun/vac/algorithms/sorting/QuickSort.hpp>
#include <fun/vac/algorithms/sorting/SelectionSort.hpp>

namespace Sorting {

    template <typename T>
    void bubbleSort(std::vector<T> &arr) {
        BubbleSort::sort(arr);
    }

    template <typename T>
    void insertionSort(std::vector<T> &arr) {
        InsertionSort::sort(arr);
    }

    template <typename T>
    void mergeSort(std::vector<T> &arr) {
        MergeSort::sort(arr);
    }

    template <typename T>
    void quickSort(std::vector<T> &arr) {
        QuickSort::sort(arr);
    }

    template <typename T>
    void selectionSort(std::vector<T> &arr) {
        SelectionSort::sort(arr);
    }
}

#endif
