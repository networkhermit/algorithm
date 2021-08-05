<?php
    namespace muse\util\TestRunner;

    function pick(callable $func): void {
        static $TestRunnerItemIndex = 0;

        if ($func()) {
            printf("✓ Item [%d] PASSED\n", $TestRunnerItemIndex);
        } else {
            fprintf(STDERR, "✗ Item [%d] FAILED\n", $TestRunnerItemIndex);
        }

        $TestRunnerItemIndex++;
    }
?>
