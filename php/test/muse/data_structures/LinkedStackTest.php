<?php
    require_once "muse/data_structures/LinkedStack.php";
    require_once "muse/util/TestRunner.php";

    use muse\data_structures\LinkedStack\LinkedStack;
    use muse\util\TestRunner;

    function testLinkedStack(): bool {
        $size = 8192;

        $stack = new LinkedStack();

        for ($i = 1; $i <= $size; $i++) {
            $stack->push($i);
        }

        if ($stack->size() != $size) {
            return false;
        }

        for ($i = $size; $i > 0; $i--) {
            if ($stack->peek() != $i) {
                return false;
            }
            $stack->pop();
        }

        return $stack->isEmpty();
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testLinkedStack());
    }
?>
