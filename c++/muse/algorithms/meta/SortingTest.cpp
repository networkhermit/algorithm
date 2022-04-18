#include <muse/algorithms/meta/Sorting.hpp>
#include <muse/util/SequenceBuilder.hpp>
#include <muse/util/Sequences.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testBubbleSort() {
  size_t size = 32768;

  vector<int> arr(size);
  SequenceBuilder::packRandom(arr);

  int checksum = Sequences::parityChecksum(arr);

  Sorting::bubbleSort(arr);

  if (Sequences::parityChecksum(arr) != checksum) {
    return false;
  }

  return Sequences::isSorted(arr);
}

bool testInsertionSort() {
  size_t size = 32768;

  vector<int> arr(size);
  SequenceBuilder::packRandom(arr);

  int checksum = Sequences::parityChecksum(arr);

  Sorting::insertionSort(arr);

  if (Sequences::parityChecksum(arr) != checksum) {
    return false;
  }

  return Sequences::isSorted(arr);
}

bool testMergeSort() {
  size_t size = 32768;

  vector<int> arr(size);
  SequenceBuilder::packRandom(arr);

  int checksum = Sequences::parityChecksum(arr);

  Sorting::mergeSort(arr);

  if (Sequences::parityChecksum(arr) != checksum) {
    return false;
  }

  return Sequences::isSorted(arr);
}

bool testQuickSort() {
  size_t size = 32768;

  vector<int> arr(size);
  SequenceBuilder::packRandom(arr);

  int checksum = Sequences::parityChecksum(arr);

  Sorting::quickSort(arr);

  if (Sequences::parityChecksum(arr) != checksum) {
    return false;
  }

  return Sequences::isSorted(arr);
}

bool testSelectionSort() {
  size_t size = 32768;

  vector<int> arr(size);
  SequenceBuilder::packRandom(arr);

  int checksum = Sequences::parityChecksum(arr);

  Sorting::selectionSort(arr);

  if (Sequences::parityChecksum(arr) != checksum) {
    return false;
  }

  return Sequences::isSorted(arr);
}

int main() {
  TestRunner::pick(&testBubbleSort);

  TestRunner::pick(&testInsertionSort);

  TestRunner::pick(&testMergeSort);

  TestRunner::pick(&testQuickSort);

  TestRunner::pick(&testSelectionSort);
}
