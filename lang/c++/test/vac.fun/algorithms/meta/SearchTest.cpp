#include <vac.fun/algorithms/meta/Search.hpp>
#include <vac.fun/util/SequenceBuilder.hpp>
#include <vac.fun/util/TestRunner.hpp>

using namespace std;

bool testBinarySearch() {
    size_t size = 32768;

    vector<int> arr(size);
    SequenceBuilder::packIncreasing(arr);

    if (Search::binarySearch(arr, -1) != size) {
        return false;
    }

    if (Search::binarySearch(arr, 2'147'483'647) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (Search::binarySearch(arr, arr[i]) != i) {
            return false;
        }
    }

    return true;
}

bool testLinearSearch() {
    size_t size = 32768;

    vector<int> arr(size);
    SequenceBuilder::packIncreasing(arr);

    if (Search::linearSearch(arr, -1) != size) {
        return false;
    }

    if (Search::linearSearch(arr, 2'147'483'647) != size) {
        return false;
    }

    for (size_t i = 0; i < size; i++) {
        if (Search::linearSearch(arr, arr[i]) != i) {
            return false;
        }
    }

    return true;
}

int main() {
    TestRunner::parseTest(testBinarySearch());

    TestRunner::parseTest(testLinearSearch());
}
