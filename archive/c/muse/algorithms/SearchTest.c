#include <muse/algorithms/Search.h>
#include <muse/util/SequenceBuilder.h>
#include <muse/util/TestRunner.h>

bool testBinarySearch(void) {
  size_t size = 32768;

  int arr[size];
  SequenceBuilder_packIncreasing(arr, size);

  if (Search_binarySearch(arr, size, -1) != size) {
    return false;
  }

  if (Search_binarySearch(arr, size, 2'147'483'647) != size) {
    return false;
  }

  for (size_t i = 0; i < size; i++) {
    if (Search_binarySearch(arr, size, arr[i]) != i) {
      return false;
    }
  }

  return true;
}

bool testLinearSearch(void) {
  size_t size = 32768;

  int arr[size];
  SequenceBuilder_packIncreasing(arr, size);

  if (Search_linearSearch(arr, size, -1) != size) {
    return false;
  }

  if (Search_linearSearch(arr, size, 2'147'483'647) != size) {
    return false;
  }

  for (size_t i = 0; i < size; i++) {
    if (Search_linearSearch(arr, size, arr[i]) != i) {
      return false;
    }
  }

  return true;
}

int main(void) {
  TestRunner_pick(&testBinarySearch);

  TestRunner_pick(&testLinearSearch);
}
