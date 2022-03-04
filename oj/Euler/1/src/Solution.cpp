#include <iostream>

using namespace std;

long bruteForce(int sentinel) {
    long sum = 0;
    for (int i = 0; i < sentinel; i++) {
        if (i % 3 == 0 || i % 5 == 0) {
            sum += static_cast<long>(i);
        }
    }
    return sum;
}

long arithmeticSeries(int sentinel, int unit) {
    long n = static_cast<long>((sentinel - 1) / unit);

    return static_cast<long>(unit) * (n * (1 + n) >> 1);
}

long process(int sentinel) {
    long sum = 0;

    sum += arithmeticSeries(sentinel, 3);
    sum -= arithmeticSeries(sentinel, 15);
    sum += arithmeticSeries(sentinel, 5);

    return sum;
}

int main() {
    int testCases;

    int sentinel;

    while (cin >> testCases) {
        for (int i = 0; i < testCases; i++) {
            cin >> sentinel;
            cout << process(sentinel) << endl;
        }
    }
}
