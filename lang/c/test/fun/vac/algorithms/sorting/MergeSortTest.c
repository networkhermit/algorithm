#include <fun/vac/algorithms/sorting/MergeSort.h>
#include <fun/vac/util/SequenceBuilder.h>
#include <fun/vac/util/Sequences.h>
#include <fun/vac/util/TestRunner.h>

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

int main(void) {
    TestRunner_parseTest(testMergeSort());
}
