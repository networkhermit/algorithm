#include <fun/vac/algorithms/sorting/QuickSort.h>
#include <fun/vac/util/SequenceBuilder.h>
#include <fun/vac/util/Sequences.h>
#include <fun/vac/util/TestRunner.h>

bool testQuickSort(void) {
    size_t size = 32768;

    int arr[size];
    SequenceBuilder_packRandom(arr, size);

    int checksum = Sequences_parityChecksum(arr, size);

    QuickSort_sort(arr, size);

    if (Sequences_parityChecksum(arr, size) != checksum) {
        return false;
    }

    if (!Sequences_isSorted(arr, size)) {
        return false;
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testQuickSort());
}
