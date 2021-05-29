#include <vac.fun/algorithms/sorting/SelectionSort.h>
#include <vac.fun/util/SequenceBuilder.h>
#include <vac.fun/util/Sequences.h>
#include <vac.fun/util/TestRunner.h>

bool testSelectionSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    SelectionSort_sort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    return Sequences_isSorted(arr, size);
}

int main(void) {
    TestRunner_parseTest(testSelectionSort());
}
