<?php
    require_once "vac.fun/util/TestRunner.php";

    use vac\fun\util\TestRunner;

    function testParseTest(): void {
        for ($i = 0; $i < 10; $i++) {
            TestRunner\parseTest(($i & 1) != 0);
        }
    }

    if (count(debug_backtrace()) == 0) {
        testParseTest();
    }
?>
