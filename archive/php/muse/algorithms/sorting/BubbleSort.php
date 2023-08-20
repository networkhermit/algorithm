<?php

declare(strict_types=1);

namespace muse\algorithms\sorting\BubbleSort;

function sort(array &$arr): void
{
    $margin = 0;
    $unsorted = count($arr);

    while ($unsorted > 1) {
        $margin = 0;
        for ($i = 1; $i < $unsorted; $i++) {
            if ($arr[$i - 1] > $arr[$i]) {
                [$arr[$i - 1], $arr[$i]] = [$arr[$i], $arr[$i - 1]];
                $margin = $i;
            }
        }
        $unsorted = $margin;
    }
}
