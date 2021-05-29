<?php
    namespace vac\fun\util\RandomFactory;

    function seed(): void {
        // preserve for consistent interface
    }

    function integerN(int $min, int $max): int {
        return random_int($min, $max);
    }

    function generateInteger(): int {
        return integerN(0, 2147483647);
    }

    function generateEven(): int {
        return generateInteger() >> 1 << 1;
    }

    function generateOdd(): int {
        return generateInteger() | 1;
    }
?>
