#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

bool *bruteForce(int n) {
    bool *arr = (bool *) malloc(n * sizeof(bool));

    memset(arr, false, n);

    for (int i = 1; i <= n; i++) {
        for (int j = i - 1; j < n; j += i) {
            arr[j] = !arr[j];
        }
    }

    return arr;
}

bool *process(int n) {
    bool *arr = (bool *) malloc(n * sizeof(bool));

    int temp;
    for (int i = 1; i <= n; i++) {
        temp = (int) sqrt((double) i);
        arr[i - 1] = temp * temp == i;
    }

    return arr;
}

int main(void) {
    const int N = 100;

    bool *arr = process(N);

    for (int i = 0; i < N; i++) {
        if (arr[i]) {
            printf("%d\t  OPEN\n", i + 1);
        }
    }
}
