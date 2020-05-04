<?php
    namespace fun\vac\algorithms\number_theory\PrimeSieves;

    function sieveOfEratosthenes(int $n): array {
        if ($n < 2) {
            return array();
        }

        $size = ($n + 1) >> 1;

        $arr = array_pad(array(), $size, false);

        $numPrimes = $size;
        for ($i = 3, $bound = (int) sqrt($n) + 1; $i < $bound; $i += 2) {
            if ($arr[$i >> 1]) {
                continue;
            }
            for ($j = $i * $i; $j <= $n; $j += $i << 1) {
                if (!$arr[$j >> 1]) {
                    $arr[$j >> 1] = true;
                    $numPrimes--;
                }
            }
        }

        $primes = array_pad(array(), $numPrimes, 0);

        $primes[0] = 2;

        $curr = 1;
        for ($i = 3, $bound = $n + 1; $i < $bound; $i += 2) {
            if (!$arr[$i >> 1]) {
                $primes[$curr] = $i;
                $curr++;
            }
        }

        return $primes;
    }
?>
