<?php
    namespace muse\algorithms\number_theory\Coprimality;

    require_once "muse/algorithms/number_theory/GreatestCommonDivisor.php";

    use muse\algorithms\number_theory\GreatestCommonDivisor;

    function reduceToBinaryGCD(int $m, int $n): bool {
        return GreatestCommonDivisor\iterativeBinaryGCD($m, $n) == 1;
    }

    function reduceToEuclidean(int $m, int $n): bool {
        return GreatestCommonDivisor\iterativeEuclidean($m, $n) == 1;
    }
?>
