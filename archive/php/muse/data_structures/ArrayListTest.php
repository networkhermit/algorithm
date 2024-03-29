<?php

declare(strict_types=1);

require_once "muse/data_structures/ArrayList.php";
require_once "muse/util/TestRunner.php";

use muse\data_structures\ArrayList;
use muse\util\TestRunner;

function testArrayList(): bool
{
    $size = 8192;

    $list = new ArrayList(0);

    for ($i = 1; $i <= $size; $i++) {
        $list->append($i);
    }

    $list->shrink();

    if ($list->size() != $size) {
        return false;
    }

    if ($list->capacity() != $size) {
        return false;
    }

    for ($i = 0; $i < $size; $i++) {
        if ($list->get($i) != $i + 1) {
            return false;
        }
    }

    for ($i = 0; $i < $size; $i++) {
        $list->set($i, $size - $i);
    }

    for ($i = 0; $i < $size; $i++) {
        if ($list->get($i) != $size - $i) {
            return false;
        }
    }

    for ($i = 0, $j = $size - 1; $i < $j; $i++, $j--) {
        $x = $list->get($i);
        $y = $list->get($j);

        $list->remove($i);
        $list->insert($i, $y);
        $list->remove($j);
        $list->insert($j, $x);
    }

    for ($i = 0; $i < $size; $i++) {
        if ($list->get($i) != $i + 1) {
            return false;
        }
    }

    for ($i = $size; $i >= 1; $i--) {
        if ($list->back() != $i) {
            return false;
        }
        $list->eject();
    }

    $list->shrink();

    if (!$list->isEmpty()) {
        return false;
    }

    return $list->capacity() == 0;
}

if (count(debug_backtrace()) == 0) {
    TestRunner\pick("testArrayList");
}
