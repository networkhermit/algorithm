<?php
    namespace fun\vac\util\RandomFactory;

    function launch(): void {
        mt_srand(time());
    }

    function integerN(int $min, int $max): int {
        return $min + mt_rand() % ($max - $min);
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
