#include <fun/vac/algorithms/number_theory/Factorial.h>
#include <fun/vac/util/TestRunner.h>

bool testFactorial(void) {
    long mapping[][2] = {
        { 0,         1},
        { 1,         1},
        { 2,         2},
        { 3,         6},
        { 4,        24},
        { 5,       120},
        { 6,       720},
        { 7,      5040},
        { 8,     40320},
        { 9,    362880},
        {10,   3628800},
        {11,  39916800},
        {12, 479001600},
    };

    size_t instances = sizeof(mapping) / sizeof(mapping[0]);

    for (size_t i = 0; i < instances; i++) {
        if (Factorial_iterativeProcedure(mapping[i][0]) != mapping[i][1]) {
            return false;
        }
    }

    for (size_t i = 0; i < instances; i++) {
        if (Factorial_recursiveProcedure(mapping[i][0]) != mapping[i][1]) {
            return false;
        }
    }

    return true;
}

int main(void) {
    TestRunner_parseTest(testFactorial());
}
