#ifndef VAC_FUN_ALGORITHMS_META_SORTING_H
#define VAC_FUN_ALGORITHMS_META_SORTING_H 1

#ifndef SORTING_TYPE
typedef int sorting_type;
#else
typedef SORTING_TYPE sorting_type;
// clang-format off
#define BUBBLE_SORT_TYPE    SORTING_TYPE
#define INSERTION_SORT_TYPE SORTING_TYPE
#define MERGE_SORT_TYPE     SORTING_TYPE
#define QUICK_SORT_TYPE     SORTING_TYPE
#define SELECTION_SORT_TYPE SORTING_TYPE
// clang-format on
#endif

#include <vac.fun/algorithms/sorting/BubbleSort.h>
#include <vac.fun/algorithms/sorting/InsertionSort.h>
#include <vac.fun/algorithms/sorting/MergeSort.h>
#include <vac.fun/algorithms/sorting/QuickSort.h>
#include <vac.fun/algorithms/sorting/SelectionSort.h>

void Sorting_bubbleSort(sorting_type *arr, size_t length) {
    BubbleSort_sort(arr, length);
}

void Sorting_insertionSort(sorting_type *arr, size_t length) {
    InsertionSort_sort(arr, length);
}

void Sorting_mergeSort(sorting_type *arr, size_t length) {
    MergeSort_sort(arr, length);
}

void Sorting_quickSort(sorting_type *arr, size_t length) {
    QuickSort_sort(arr, length);
}

void Sorting_selectionSort(sorting_type *arr, size_t length) {
    SelectionSort_sort(arr, length);
}

#endif