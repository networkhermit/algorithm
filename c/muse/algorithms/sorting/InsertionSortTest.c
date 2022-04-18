#include <muse/algorithms/sorting/InsertionSort.h>
#include <muse/util/SequenceBuilder.h>
#include <muse/util/Sequences.h>
#include <muse/util/TestRunner.h>

bool testInsertionSort(void) {
  size_t size = 32768;

  int arr[size];
  SequenceBuilder_packRandom(arr, size);

  int checksum = Sequences_parityChecksum(arr, size);

  InsertionSort_sort(arr, size);

  if (Sequences_parityChecksum(arr, size) != checksum) {
    return false;
  }

  return Sequences_isSorted(arr, size);
}

int main(void) { TestRunner_pick(&testInsertionSort); }
