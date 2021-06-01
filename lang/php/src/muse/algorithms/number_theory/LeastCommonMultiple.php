<?php
    namespace muse\algorithms\number_theory\LeastCommonMultiple;

    require_once "muse/algorithms/number_theory/GreatestCommonDivisor.php";

    use muse\algorithms\number_theory\GreatestCommonDivisor;

    function reduceToBinaryGCD(int $m, int $n): int {
        if ($m < 0) {
            $m = -$m;
        }
        if ($n < 0) {
            $n = -$n;
        }

        if ($m == 0 || $n == 0) {
            return 0;
        }
        return $m / GreatestCommonDivisor\iterativeBinaryGCD($m, $n) * $n;
    }

    function reduceToEuclidean(int $m, int $n): int {
        if ($m < 0) {
            $m = -$m;
        }
        if ($n < 0) {
            $n = -$n;
        }

        if ($m == 0 || $n == 0) {
            return 0;
        }
        return $m / GreatestCommonDivisor\iterativeEuclidean($m, $n) * $n;
    }
?>
