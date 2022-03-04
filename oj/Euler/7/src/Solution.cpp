#include <cmath>
#include <cstring>
#include <iostream>

using namespace std;

int draft(int n) {
    if (n == 1) {
        return 2;
    }

    int primes[n];

    primes[0] = 2;

    int k = 1;

    bool isPrime;
    int bound;

    for (int i = 3; i < 2147483647; i += 2) {
        isPrime = true;
        bound = static_cast<int>(sqrt(static_cast<double>(i)));
        for (int j = 1; j < k; j++) {
            if (primes[j] > bound) {
                break;
            }
            if (i % primes[j] == 0) {
                isPrime = false;
                break;
            }
        }
        if (isPrime) {
            primes[k] = i;
            k++;
            if (k == n) {
                break;
            }
        }
    }

    return primes[n - 1];
}

int process(int n) {
    if (n == 1) {
        return 2;
    }

    int ceiling;

    if (n < 6) {
        ceiling = (n << 1) + 1;
    } else {
        double cache = log(static_cast<double>(n));
        ceiling = static_cast<int>(static_cast<double>(n) * (cache + log(cache)));
    }

    int size = (ceiling + 1) >> 1;

    bool arr[size];
    memset(arr, false, size);

    for (int i = 3, bound = static_cast<int>(sqrt(static_cast<double>(ceiling))) + 1; i < bound; i += 2) {
        if (arr[i >> 1]) {
            continue;
        }
        for (int j = i * i; j <= ceiling; j += i << 1) {
            if (!arr[j >> 1]) {
                arr[j >> 1] = true;
            }
        }
    }

    int k = 1;
    for (int i = 1; i < size; i++) {
        if (!arr[i]) {
            k++;
            if (k == n) {
                return (i << 1) + 1;
            }
        }
    }

    return 0;
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
