<?php

declare(strict_types=1);

namespace muse\util\RandomFactory;

function genIntN(int $min, int $max): int
{
    return random_int($min, $max);
}

function genInt(): int
{
    return genIntN(1, 2147483647);
}

function genEven(): int
{
    return genInt() >> 1 << 1;
}

function genOdd(): int
{
    return genInt() | 1;
}
