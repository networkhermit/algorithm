<?php

declare(strict_types=1);

require_once "muse/util/RandomFactory.php";
require_once "muse/util/TestRunner.php";

use muse\util\RandomFactory;
use muse\util\TestRunner;

function testGenIntN(): bool
{
    $value = 0;
    for ($i = 0; $i < 8192; $i++) {
        if (RandomFactory\genIntN(0, 0) != 0) {
            return false;
        }

        if (RandomFactory\genIntN(1, 1) != 1) {
            return false;
        }

        $value = RandomFactory\genIntN(0, 1);
        if ($value < 0 || $value > 1) {
            return false;
        }

        $value = RandomFactory\genIntN(100, 10000);
        if (RandomFactory\genIntN($value, $value) != $value) {
            return false;
        }
        if ($value < 100 || $value > 10000) {
            return false;
        }
    }

    return true;
}

function testGenEven(): bool
{
    for ($i = 0; $i < 8192; $i++) {
        if ((RandomFactory\genEven() & 1) != 0) {
            return false;
        }
    }

    return true;
}

function testGenOdd(): bool
{
    for ($i = 0; $i < 8192; $i++) {
        if ((RandomFactory\genOdd() & 1) == 0) {
            return false;
        }
    }

    return true;
}

if (count(debug_backtrace()) == 0) {
    TestRunner\pick("testGenIntN");

    TestRunner\pick("testGenEven");

    TestRunner\pick("testGenOdd");
}
