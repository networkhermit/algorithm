<?php
    require_once "fun/vac/data_structures/LinkedStack.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\data_structures\LinkedStack\LinkedStack;
    use fun\vac\util\TestRunner;

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

        if (!$stack->isEmpty()) {
            return false;
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testLinkedStack());
    }
?>
