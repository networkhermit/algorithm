#include <limits.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const char *fizzBuzz(int n) {
    const size_t MAX_DIGIT = (size_t) log10((double) INT_MAX) + 1;
    const size_t MAX_ALPHA = strlen("FizzBuzz");
    const size_t MAX_LENGTH = MAX_DIGIT > MAX_ALPHA ? MAX_DIGIT : MAX_ALPHA;

    char *word = (char *) calloc(MAX_LENGTH, sizeof(char));

    if (n % 3 == 0) {
        strncat(word, "Fizz", strlen("Fizz") + 1);
    }
    if (n % 5 == 0) {
        strncat(word, "Buzz", strlen("Buzz") + 1);
    }

    if (word[0] == '\0') {
        snprintf(word, MAX_DIGIT + 1, "%d", n);
    }

    return word;
}

int main(void) {
    for (int i = 1; i <= 100; i++) {
        printf("%s\n", fizzBuzz(i));
    }
}
