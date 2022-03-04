#include <iostream>

using namespace std;

bool draft(int n) {
    int k = 0;
    int arr[10] = {0};

    while (n != 0) {
        arr[k] = n % 10;
        n /= 10;
        k++;
    }

    for (int i = 0, j = k - 1, bound = (k >> 1) + 1; i < bound; i++, j--) {
        if (arr[i] != arr[j]) {
            return false;
        }
    }

    return true;
}

bool isPalindrome(int n) {
    int original = n;

    int reversed = 0;
    while (n != 0) {
        reversed = reversed * 10 + n % 10;
        n /= 10;
    }

    return original == reversed;
}

bool exhaustiveMatch(int n) {
    int product;

    for (int i = 999; i > 99; i--) {
        for (int j = i; j > 99; j--) {
            product = i * j;
            if (product == n) {
                return true;
            } else if (product < n) {
                break;
            }
        }
    }

    return false;
}

int process(int sentinel) {
    for (int i = sentinel - 1; i > 10000; i--) {
        if (isPalindrome(i) && exhaustiveMatch(i)) {
            return i;
        }
    }
    return 0;
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
