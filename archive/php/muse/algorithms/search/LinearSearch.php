<?php

declare(strict_types=1);

namespace muse\algorithms\search\LinearSearch;

function find(array &$arr, $key): int
{
    foreach ($arr as $i => $v) {
        if ($v == $key) {
            return $i;
        }
    }

    return count($arr);
}
