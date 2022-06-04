<?php

declare(strict_types=1);

require_once "muse/algorithms/Sorting.php";
require_once "muse/util/SequenceBuilder.php";
require_once "muse/util/Sequences.php";
require_once "muse/util/TestRunner.php";

use muse\algorithms\Sorting;
use muse\util\SequenceBuilder;
use muse\util\Sequences;
use muse\util\TestRunner;

function testBubbleSort(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packRandom($arr);

    $checksum = Sequences\parityChecksum($arr);

    Sorting\bubbleSort($arr);

    if (Sequences\parityChecksum($arr) != $checksum) {
        return false;
    }

    return Sequences\isSorted($arr);
}

function testInsertionSort(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packRandom($arr);

    $checksum = Sequences\parityChecksum($arr);

    Sorting\insertionSort($arr);

    if (Sequences\parityChecksum($arr) != $checksum) {
        return false;
    }

    return Sequences\isSorted($arr);
}

function testMergeSort(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packRandom($arr);

    $checksum = Sequences\parityChecksum($arr);

    Sorting\mergeSort($arr);

    if (Sequences\parityChecksum($arr) != $checksum) {
        return false;
    }

    return Sequences\isSorted($arr);
}

function testQuickSort(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packRandom($arr);

    $checksum = Sequences\parityChecksum($arr);

    Sorting\quickSort($arr);

    if (Sequences\parityChecksum($arr) != $checksum) {
        return false;
    }

    return Sequences\isSorted($arr);
}

function testSelectionSort(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packRandom($arr);

    $checksum = Sequences\parityChecksum($arr);

    Sorting\selectionSort($arr);

    if (Sequences\parityChecksum($arr) != $checksum) {
        return false;
    }

    return Sequences\isSorted($arr);
}

if (count(debug_backtrace()) == 0) {
    TestRunner\pick("testBubbleSort");

    TestRunner\pick("testInsertionSort");

    TestRunner\pick("testMergeSort");

    TestRunner\pick("testQuickSort");

    TestRunner\pick("testSelectionSort");
}
