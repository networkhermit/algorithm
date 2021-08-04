<?php
    require_once "muse/util/TestRunner.php";

    use muse\util\TestRunner;

    function testParseTest(): void {
        for ($i = 0; $i < 10; $i++) {
            TestRunner\parseTest(($i & 1) != 0);
        }
    }

    if (count(debug_backtrace()) == 0) {
        testParseTest();
    }
?>
