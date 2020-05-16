<?php
    require_once "fun/vac/data_structures/ArrayQueue.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\data_structures\ArrayQueue\ArrayQueue;
    use fun\vac\util\TestRunner;

    function testArrayQueue(): bool {
        $size = 8192;

        $queue = new ArrayQueue(0);

        for ($i = 1; $i <= $size; $i++) {
            $queue->enqueue($i);
        }

        $queue->shrink();

        if ($queue->size() != $size) {
            return false;
        }

        if ($queue->capacity() != $size) {
            return false;
        }

        for ($i = 1; $i <= $size; $i++) {
            if ($queue->peek() != $i) {
                return false;
            }
            $queue->dequeue();
        }

        $queue->shrink();

        if (!$queue->isEmpty()) {
            return false;
        }

        return $queue->capacity() == 0;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testArrayQueue());
    }
?>
