<?php
    namespace muse\util\RandomFactory;

    function seed(): void {
        // preserve for consistent interface
    }

    function genIntN(int $min, int $max): int {
        return random_int($min, $max);
    }

    function genInt(): int {
        return genIntN(0, 2147483647);
    }

    function genEven(): int {
        return genInt() >> 1 << 1;
    }

    function genOdd(): int {
        return genInt() | 1;
    }
?>
