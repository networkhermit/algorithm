<?php

declare(strict_types=1);

require_once "muse/algorithms/sorting/QuickSort.php";
require_once "muse/util/SequenceBuilder.php";
require_once "muse/util/Sequences.php";
require_once "muse/util/TestRunner.php";

use muse\algorithms\sorting\QuickSort;
use muse\util\SequenceBuilder;
use muse\util\Sequences;
use muse\util\TestRunner;

function testQuickSort(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packRandom($arr);

    $checksum = Sequences\parityChecksum($arr);

    QuickSort\sort($arr);

    if (Sequences\parityChecksum($arr) != $checksum) {
        return false;
    }

    return Sequences\isSorted($arr);
}

if (count(debug_backtrace()) == 0) {
    TestRunner\pick("testQuickSort");
}
