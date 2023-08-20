<?php

declare(strict_types=1);

require_once "muse/algorithms/number_theory/Factorial.php";
require_once "muse/util/TestRunner.php";

use muse\algorithms\number_theory\Factorial;
use muse\util\TestRunner;

function testFactorial(): bool
{
    $sample = [
        [0, 1],
        [1, 1],
        [2, 2],
        [3, 6],
        [4, 24],
        [5, 120],
        [6, 720],
        [7, 5040],
        [8, 40320],
        [9, 362880],
        [10, 3_628_800],
        [11, 39_916_800],
        [12, 479_001_600],
    ];

    for ($i = 0, $size = count($sample); $i < $size; $i++) {
        if (Factorial\iterativeProcedure($sample[$i][0]) != $sample[$i][1]) {
            return false;
        }
    }

    for ($i = 0, $size = count($sample); $i < $size; $i++) {
        if (Factorial\recursiveProcedure($sample[$i][0]) != $sample[$i][1]) {
            return false;
        }
    }

    return true;
}

if (count(debug_backtrace()) == 0) {
    TestRunner\pick("testFactorial");
}
