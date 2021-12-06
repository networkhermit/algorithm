<?php
    declare(strict_types=1);

    namespace muse\algorithms\meta\NumberTheory;

    require_once "muse/algorithms/number_theory/Coprimality.php";
    require_once "muse/algorithms/number_theory/Factorial.php";
    require_once "muse/algorithms/number_theory/FibonacciNumber.php";
    require_once "muse/algorithms/number_theory/GreatestCommonDivisor.php";
    require_once "muse/algorithms/number_theory/LeastCommonMultiple.php";
    require_once "muse/algorithms/number_theory/Parity.php";
    require_once "muse/algorithms/number_theory/Primality.php";
    require_once "muse/algorithms/number_theory/PrimeSieves.php";

    use muse\algorithms\number_theory\Coprimality;
    use muse\algorithms\number_theory\Factorial;
    use muse\algorithms\number_theory\FibonacciNumber;
    use muse\algorithms\number_theory\GreatestCommonDivisor;
    use muse\algorithms\number_theory\LeastCommonMultiple;
    use muse\algorithms\number_theory\Parity;
    use muse\algorithms\number_theory\Primality;
    use muse\algorithms\number_theory\PrimeSieves;

    function isCoprime(int $m, int $n): bool {
        return Coprimality\reduceToBinaryGCD($m, $n);
    }

    function factorial(int $n): int {
        return Factorial\iterativeProcedure($n);
    }

    function fibonacci(int $n): int {
        return FibonacciNumber\iterativeProcedure($n);
    }

    function gcd(int $m, int $n): int {
        return GreatestCommonDivisor\iterativeBinaryGCD($m, $n);
    }

    function lcm(int $m, int $n): int {
        return LeastCommonMultiple\reduceToBinaryGCD($m, $n);
    }

    function isEven(int $n): bool {
        return Parity\bitwiseIsEven($n);
    }

    function isOdd(int $n): bool {
        return Parity\bitwiseIsOdd($n);
    }

    function isPrime(int $n): bool {
        return Primality\isPrime($n);
    }

    function isComposite(int $n): bool {
        return Primality\isComposite($n);
    }

    function sieveOfPrimes(int $n): array {
        return PrimeSieves\sieveOfEratosthenes($n);
    }
?>
