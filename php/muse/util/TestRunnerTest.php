<?php
    declare(strict_types=1);

    require_once "muse/util/TestRunner.php";

    use muse\util\TestRunner;

    function testPick(): void {
        for ($i = 0; $i < 10; $i++) {
            TestRunner\pick(fn() => ($i & 1) != 0);
        }
    }

    if (count(debug_backtrace()) == 0) {
        testPick();
    }
?>
