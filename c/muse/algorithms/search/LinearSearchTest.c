#include <muse/algorithms/search/LinearSearch.h>
#include <muse/util/SequenceBuilder.h>
#include <muse/util/TestRunner.h>

bool testLinearSearch(void) {
  size_t size = 32768;

  int arr[size];
  SequenceBuilder_packIncreasing(arr, size);

  if (LinearSearch_find(arr, size, -1) != size) {
    return false;
  }

  if (LinearSearch_find(arr, size, 2147483647) != size) {
    return false;
  }

  for (size_t i = 0; i < size; i++) {
    if (LinearSearch_find(arr, size, arr[i]) != i) {
      return false;
    }
  }

  return true;
}

int main(void) { TestRunner_pick(&testLinearSearch); }
