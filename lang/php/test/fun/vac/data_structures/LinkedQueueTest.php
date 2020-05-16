<?php
    require_once "fun/vac/data_structures/LinkedQueue.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\data_structures\LinkedQueue\LinkedQueue;
    use fun\vac\util\TestRunner;

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
