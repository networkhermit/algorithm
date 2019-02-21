#include <fun/vac/algorithms/sorting/BubbleSort.hpp>
#include <fun/vac/util/SequenceBuilder.hpp>
#include <fun/vac/util/Sequences.hpp>
#include <fun/vac/util/TestRunner.hpp>

using namespace std;

bool testBubbleSort() {
    size_t size = 32768;

    vector<int> arr (size);
    SequenceBuilder::packRandom(arr);

    int checksum = Sequences::parityChecksum(arr);

    BubbleSort::sort(arr);

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
}
