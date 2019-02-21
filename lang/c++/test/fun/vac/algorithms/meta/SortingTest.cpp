#include <fun/vac/algorithms/meta/Sorting.hpp>
#include <fun/vac/util/SequenceBuilder.hpp>
#include <fun/vac/util/Sequences.hpp>
#include <fun/vac/util/TestRunner.hpp>

using namespace std;

bool testBubbleSort() {
    size_t size = 32768;

    vector<int> arr (size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    Sorting::bubbleSort(arr);

    if (Sequences::parityChecksum(arr) != checksum) {
        return false;
    }

    if (!Sequences::isSorted(arr)) {
        return false;
    }

    return true;
}

bool testInsertionSort() {
    size_t size = 32768;

    vector<int> arr (size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    Sorting::insertionSort(arr);

    if (Sequences::parityChecksum(arr) != checksum) {
        return false;
    }

    if (!Sequences::isSorted(arr)) {
        return false;
    }

    return true;
}

bool testMergeSort() {
    size_t size = 32768;

    vector<int> arr (size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    Sorting::mergeSort(arr);

    if (Sequences::parityChecksum(arr) != checksum) {
        return false;
    }

    if (!Sequences::isSorted(arr)) {
        return false;
    }

    return true;
}

bool testQuickSort() {
    size_t size = 32768;

    vector<int> arr (size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    Sorting::quickSort(arr);

    if (Sequences::parityChecksum(arr) != checksum) {
        return false;
    }

    if (!Sequences::isSorted(arr)) {
        return false;
    }

    return true;
}

bool testSelectionSort() {
    size_t size = 32768;

    vector<int> arr (size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    Sorting::selectionSort(arr);

    if (Sequences::parityChecksum(arr) != checksum) {
        return false;
    }

    if (!Sequences::isSorted(arr)) {
        return false;
    }

    return true;
}

int main() {

    TestRunner::parseTest(testBubbleSort());

    TestRunner::parseTest(testInsertionSort());

    TestRunner::parseTest(testMergeSort());

    TestRunner::parseTest(testQuickSort());

    TestRunner::parseTest(testSelectionSort());
}
