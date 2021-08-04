package Factorial

import "testing"

func TestFactorial(t *testing.T) {
    sample := [][2]int64{
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

    for i := range sample {
        if IterativeProcedure(sample[i][0]) != sample[i][1] {
            t.FailNow()
        }
    }

    for i := range sample {
        if RecursiveProcedure(sample[i][0]) != sample[i][1] {
            t.FailNow()
        }
    }
}
