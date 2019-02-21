<?php
    namespace fun\vac\algorithms\number_theory\Coprimality;

    require_once "fun/vac/algorithms/number_theory/GreatestCommonDivisor.php";

    use fun\vac\algorithms\number_theory\GreatestCommonDivisor;

    function reduceToBinaryGCD(int $m, int $n): bool {
        return GreatestCommonDivisor\iterativeBinaryGCD($m, $n) == 1;
    }

    function reduceToEuclidean(int $m, int $n): bool {
        return GreatestCommonDivisor\iterativeEuclidean($m, $n) == 1;
    }
?>
