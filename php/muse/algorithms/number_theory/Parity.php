<?php
    namespace muse\algorithms\number_theory\Parity;

    function moduloIsEven(int $n): bool {
        return $n % 2 == 0;
    }

    function moduloIsOdd(int $n): bool {
        return $n % 2 != 0;
    }

    function bitwiseIsEven(int $n): bool {
        return ($n & 1) == 0;
    }

    function bitwiseIsOdd(int $n): bool {
        return ($n & 1) != 0;
    }
?>
