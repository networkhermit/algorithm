<?php

declare(strict_types=1);

require_once "muse/algorithms/Search.php";
require_once "muse/util/SequenceBuilder.php";
require_once "muse/util/TestRunner.php";

use muse\algorithms\Search;
use muse\util\SequenceBuilder;
use muse\util\TestRunner;

function testBinarySearch(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packIncreasing($arr);

    if (Search\binarySearch($arr, -1) != $size) {
        return false;
    }

    if (Search\binarySearch($arr, 2_147_483_647) != $size) {
        return false;
    }

    foreach ($arr as $i => $v) {
        if (Search\binarySearch($arr, $v) != $i) {
            return false;
        }
    }

    return true;
}

function testLinearSearch(): bool
{
    $size = 32768;

    $arr = array_pad(array(), $size, 0);
    SequenceBuilder\packIncreasing($arr);

    if (Search\linearSearch($arr, -1) != $size) {
        return false;
    }

    if (Search\linearSearch($arr, 2_147_483_647) != $size) {
        return false;
    }

    foreach ($arr as $i => $v) {
        if (Search\linearSearch($arr, $v) != $i) {
            return false;
        }
    }

    return true;
}

if (count(debug_backtrace()) == 0) {
    TestRunner\pick("testBinarySearch");

    TestRunner\pick("testLinearSearch");
}
