#include <muse/algorithms/sorting/InsertionSort.hpp>
#include <muse/util/SequenceBuilder.hpp>
#include <muse/util/Sequences.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testInsertionSort() {
  size_t size = 32768;

  vector<int> arr(size);
  SequenceBuilder::packRandom(arr);

  int checksum = Sequences::parityChecksum(arr);

  InsertionSort::sort(arr);

  if (Sequences::parityChecksum(arr) != checksum) {
    return false;
  }

  return Sequences::isSorted(arr);
}

int main() { TestRunner::pick(&testInsertionSort); }
