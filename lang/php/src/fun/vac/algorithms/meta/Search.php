<?php
    namespace fun\vac\algorithms\meta\Search;

    require_once "fun/vac/algorithms/search/BinarySearch.php";
    require_once "fun/vac/algorithms/search/LinearSearch.php";

    use fun\vac\algorithms\search\BinarySearch;
    use fun\vac\algorithms\search\LinearSearch;

    function binarySearch(array &$arr, $key): int {
        return BinarySearch\find($arr, $key);
    }

    function linearSearch(array &$arr, $key): int {
        return LinearSearch\find($arr, $key);
    }
?>
