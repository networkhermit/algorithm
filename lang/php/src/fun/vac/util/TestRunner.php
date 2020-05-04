<?php
    namespace fun\vac\util\TestRunner;

    function parseTest(bool $ok): void {
        static $TestRunnerItemIndex = 0;

        if ($ok) {
            printf("✓ Item [%d] PASSED\n", $TestRunnerItemIndex);
        } else {
            fprintf(STDERR, "✗ Item [%d] FAILED\n", $TestRunnerItemIndex);
        }

        $TestRunnerItemIndex++;
    }
?>
