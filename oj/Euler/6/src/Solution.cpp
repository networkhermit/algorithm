#include <iostream>

using namespace std;

long draft(int n) {
    long v = static_cast<long>(n);
    long result = 0;

    for (long i = 1; i < v; i++) {
        result += i * (v - i) * (i + 1 + v);
    }

    return result;
}

long process(int n) {
    long v = static_cast<long>(n);
    return (v * (v + 1) >> 1) * (v - 1) / 3 * (3 * v + 2) >> 1;
}

int main() {
    int testCases;

    int n;

    while (cin >> testCases) {
        for (int i = 0; i < testCases; i++) {
            cin >> n;
            cout << process(n) << endl;
        }
    }
}
