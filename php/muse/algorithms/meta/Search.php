<?php

declare(strict_types=1);

namespace muse\algorithms\meta\Search;

require_once "muse/algorithms/search/BinarySearch.php";
require_once "muse/algorithms/search/LinearSearch.php";

use muse\algorithms\search\BinarySearch;
use muse\algorithms\search\LinearSearch;

function binarySearch(array &$arr, $key): int
{
    return BinarySearch\find($arr, $key);
}

function linearSearch(array &$arr, $key): int
{
    return LinearSearch\find($arr, $key);
}
