#include <muse/algorithms/sorting/MergeSort.h>
#include <muse/util/SequenceBuilder.h>
#include <muse/util/Sequences.h>
#include <muse/util/TestRunner.h>

bool testMergeSort(void) {
  size_t size = 32768;

  int arr[size];
  SequenceBuilder_packRandom(arr, size);

  int checksum = Sequences_parityChecksum(arr, size);

  MergeSort_sort(arr, size);

  if (Sequences_parityChecksum(arr, size) != checksum) {
    return false;
  }

  return Sequences_isSorted(arr, size);
}

int main(void) { TestRunner_pick(&testMergeSort); }
