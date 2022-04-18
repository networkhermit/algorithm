<?php

declare(strict_types=1);

namespace muse\algorithms\number_theory\Primality;

function isPrime(int $n): bool
{
    if ($n < 2) {
        return false;
    }
    if (($n & 1) == 0 && $n != 2) {
        return false;
    }

    for ($i = 3, $bound = (int) sqrt($n) + 1; $i < $bound; $i += 2) {
        if ($n % $i == 0) {
            return false;
        }
    }

    return true;
}

function isComposite(int $n): bool
{
    return $n > 1 && !isPrime($n);
}
