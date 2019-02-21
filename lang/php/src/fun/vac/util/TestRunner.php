<?php
    namespace fun\vac\util\TestRunner;

    function parseTest(bool $ok): void {
        static $TestRunner_TestIndex = 0;

        if ($ok) {
            printf("< Test [%d] Passed >\n", $TestRunner_TestIndex);
        } else {
            fprintf(STDERR, "X Test [%d] Failed X\n", $TestRunner_TestIndex);
        }

        $TestRunner_TestIndex += 1;
    }
?>
