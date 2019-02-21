<?php
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\util\TestRunner;

    function testParseTest(): void {
        for ($i = 0; $i < 10; $i++) {
            if (($i & 1) == 0) {
                TestRunner\parseTest(false);
            } else {
                TestRunner\parseTest(true);
            }
        }
    }

    if (count(debug_backtrace()) == 0) {
        testParseTest();
    }
?>
