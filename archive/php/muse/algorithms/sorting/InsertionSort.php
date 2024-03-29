<?php

declare(strict_types=1);

namespace muse\algorithms\sorting\InsertionSort;

function sort(array &$arr): void
{
    $target = null;

    $cursor = 0;

    for ($i = 1, $length = count($arr); $i < $length; $i++) {
        $target = $arr[$i];
        for ($cursor = $i; $cursor > 0; $cursor--) {
            if ($arr[$cursor - 1] <= $target) {
                break;
            }
            $arr[$cursor] = $arr[$cursor - 1];
        }
        $arr[$cursor] = $target;
    }
}
