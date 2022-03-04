#include <iostream>

using namespace std;

long process(long sentinel) {
    long sum = 0;

    long prev = 0;
    long curr = 2;

    long temp;
    while (curr < sentinel) {
        sum += curr;
        temp = prev + (curr << 2);
        prev = curr;
        curr = temp;
    }

    return sum;
}

int main() {
    int testCases;

    long sentinel;

    while (cin >> testCases) {
        for (int i = 0; i < testCases; i++) {
            cin >> sentinel;
            cout << process(sentinel) << endl;
        }
    }
}
