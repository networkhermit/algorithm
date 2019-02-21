<?php
    namespace fun\vac\algorithms\number_theory\LeastCommonMultiple;

    require_once "fun/vac/algorithms/number_theory/GreatestCommonDivisor.php";

    use fun\vac\algorithms\number_theory\GreatestCommonDivisor;

    function reduceToBinaryGCD(int $m, int $n): int {
        if ($m < 0) {
            $m = -$m;
        }
        if ($n < 0) {
            $n = -$n;
        }

        if ($m == 0 || $n == 0) {
            return 0;
        } else {
            return $m / GreatestCommonDivisor\iterativeBinaryGCD($m, $n) * $n;
        }
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
        } else {
            return $m / GreatestCommonDivisor\iterativeEuclidean($m, $n) * $n;
        }
    }
?>
