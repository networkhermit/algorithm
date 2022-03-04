#include <cmath>
#include <cstring>
#include <iostream>

using namespace std;

long process(int n) {
    if (n < 2) {
        return 1;
    }

    int size = (n + 1) >> 1;

    bool sievedFactors[size];
    memset(sievedFactors, false, size);

    for (int i = 3, bound = static_cast<int>(sqrt(static_cast<double>(n))) + 1; i < bound; i += 2) {
        if (sievedFactors[i >> 1]) {
            continue;
        }
        for (int j = i * i; j <= n; j += i << 1) {
            if (!sievedFactors[j >> 1]) {
                sievedFactors[j >> 1] = true;
            }
        }
    }

    long result = 1;

    for (int i = 2; i <= n; i <<= 1) {
        result <<= 1;
    }

    int factor;
    for (int i = 1; i < size; i++) {
        if (!sievedFactors[i]) {
            factor = (i << 1) + 1;
            for (int j = factor; j <= n; j *= factor) {
                result *= static_cast<long>(factor);
            }
        }
    }

    return result;
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
