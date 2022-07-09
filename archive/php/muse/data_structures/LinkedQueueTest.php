<?php

declare(strict_types=1);

require_once "muse/data_structures/LinkedQueue.php";
require_once "muse/util/TestRunner.php";

use muse\data_structures\LinkedQueue;
use muse\util\TestRunner;

function testLinkedQueue(): bool
{
    $size = 8192;

    $queue = new LinkedQueue();

    for ($i = 1; $i <= $size; $i++) {
        $queue->enqueue($i);
    }

    if ($queue->size() != $size) {
        return false;
    }

    for ($i = 1; $i <= $size; $i++) {
        if ($queue->peek() != $i) {
            return false;
        }
        $queue->dequeue();
    }

    return $queue->isEmpty();
}

if (count(debug_backtrace()) == 0) {
    TestRunner\pick("testLinkedQueue");
}
