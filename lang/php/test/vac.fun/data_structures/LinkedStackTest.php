<?php
    require_once "vac.fun/data_structures/LinkedStack.php";
    require_once "vac.fun/util/TestRunner.php";

    use vac\fun\data_structures\LinkedStack\LinkedStack;
    use vac\fun\util\TestRunner;

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
