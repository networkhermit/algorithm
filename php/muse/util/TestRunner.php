<?php
    declare(strict_types=1);

    namespace muse\util\TestRunner;

    function pick(callable $func): void {
        static $itemIndex = 0;

        if ($func()) {
            printf("✓ Item [%d] PASSED\n", $itemIndex);
        } else {
            fprintf(STDERR, "✗ Item [%d] FAILED\n", $itemIndex);
        }

        $itemIndex++;
    }
?>
