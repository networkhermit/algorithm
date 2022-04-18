<?php

declare(strict_types=1);

function fizzBuzz(int $n): string
{
    $word = "";

    if ($n % 3 == 0) {
        $word .= "Fizz";
    }
    if ($n % 5 == 0) {
        $word .= "Buzz";
    }

    if (strlen($word) == 0) {
        $word = (string) $n;
    }

    return $word;
}

if (count(debug_backtrace()) == 0) {
    for ($i = 1; $i <= 100; $i++) {
        printf("%s\n", fizzBuzz($i));
    }
}
