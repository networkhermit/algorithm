<?php
    namespace vac\fun\algorithms\meta\Search;

    require_once "vac.fun/algorithms/search/BinarySearch.php";
    require_once "vac.fun/algorithms/search/LinearSearch.php";

    use vac\fun\algorithms\search\BinarySearch;
    use vac\fun\algorithms\search\LinearSearch;

    function binarySearch(array &$arr, $key): int {
        return BinarySearch\find($arr, $key);
    }

    function linearSearch(array &$arr, $key): int {
        return LinearSearch\find($arr, $key);
    }
?>
