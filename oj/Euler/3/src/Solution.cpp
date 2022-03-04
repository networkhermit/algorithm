#include <cmath>
#include <iostream>

using namespace std;

long process(long n) {
    if (n < 2) {
        return 1;
    }

    while ((n & 1) == 0) {
        n >>= 1;
    }

    if (n == 1) {
        return 2;
    }

    for (long i = 3, bound = static_cast<long>(sqrt(static_cast<double>(n))) + 1; i < bound; i += 2) {
        while (n % i == 0) {
            if (n == i) {
                return n;
            }
            n /= i;
        }
    }

    return n;
}

int main() {
    int testCases;

    long n;

    while (cin >> testCases) {
        for (int i = 0; i < testCases; i++) {
            cin >> n;
            cout << process(n) << endl;
        }
    }
}
