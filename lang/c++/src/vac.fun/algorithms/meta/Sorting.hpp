#ifndef VAC_FUN_ALGORITHMS_META_SORTING_HPP
#define VAC_FUN_ALGORITHMS_META_SORTING_HPP 1

#include <vac.fun/algorithms/sorting/BubbleSort.hpp>
#include <vac.fun/algorithms/sorting/InsertionSort.hpp>
#include <vac.fun/algorithms/sorting/MergeSort.hpp>
#include <vac.fun/algorithms/sorting/QuickSort.hpp>
#include <vac.fun/algorithms/sorting/SelectionSort.hpp>

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
