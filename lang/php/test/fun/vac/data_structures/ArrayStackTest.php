<?php
    require_once "fun/vac/data_structures/ArrayStack.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\data_structures\ArrayStack\ArrayStack;
    use fun\vac\util\TestRunner;

    function testArrayStack(): bool {
        $size = 8192;

        $stack = new ArrayStack(0);

        for ($i = 1; $i <= $size; $i++) {
            $stack->push($i);
        }

        $stack->shrink();

        if ($stack->size() != $size) {
            return false;
        }

        if ($stack->capacity() != $size) {
            return false;
        }

        for ($i = $size; $i > 0; $i--) {
            if ($stack->peek() != $i) {
                return false;
            }
            $stack->pop();
        }

        $stack->shrink();

        if (!$stack->isEmpty()) {
            return false;
        }

        if ($stack->capacity() != 0) {
            return false;
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testArrayStack());
    }
?>
