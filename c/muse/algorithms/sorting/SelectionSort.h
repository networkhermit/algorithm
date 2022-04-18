#ifndef MUSE_ALGORITHMS_SORTING_SELECTION_SORT_H
#define MUSE_ALGORITHMS_SORTING_SELECTION_SORT_H

#include <stddef.h>

#ifndef SELECTION_SORT_TYPE
typedef int selection_sort_type;
#else
typedef SELECTION_SORT_TYPE selection_sort_type;
#endif

void SelectionSort_sort(selection_sort_type *arr, size_t length) {
  if (length == 0) {
    return;
  }

  selection_sort_type temp;

  size_t iMin;

  for (size_t i = 0, bound = length - 1; i < bound; i++) {
    iMin = i;
    for (size_t j = i + 1; j <= bound; j++) {
      if (arr[j] < arr[iMin]) {
        iMin = j;
      }
    }
    if (iMin != i) {
      temp = arr[i];
      arr[i] = arr[iMin];
      arr[iMin] = temp;
    }
  }
}

#endif
