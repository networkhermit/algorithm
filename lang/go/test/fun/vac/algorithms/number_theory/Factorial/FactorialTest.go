package main

import (
    "fun/vac/algorithms/number_theory/Factorial"
    "fun/vac/util/TestRunner"
)

func testFactorial() bool {
    mapping := [][2]int64{
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
    }

    for i := range mapping {
        if Factorial.IterativeProcedure(mapping[i][0]) != mapping[i][1] {
            return false
        }
    }

    for i := range mapping {
        if Factorial.RecursiveProcedure(mapping[i][0]) != mapping[i][1] {
            return false
        }
    }

    return true
}

func main() {
    TestRunner.ParseTest(testFactorial())
}
