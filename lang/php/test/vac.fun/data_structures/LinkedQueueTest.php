<?php
    require_once "vac.fun/data_structures/LinkedQueue.php";
    require_once "vac.fun/util/TestRunner.php";

    use vac\fun\data_structures\LinkedQueue\LinkedQueue;
    use vac\fun\util\TestRunner;

    function testLinkedQueue(): bool {
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
        TestRunner\parseTest(testLinkedQueue());
    }
?>
