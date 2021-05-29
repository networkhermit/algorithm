<?php
    namespace vac\fun\algorithms\search\BinarySearch;

    function find(array &$arr, $key): int {
        $lo = 0;
        $hi = count($arr);

        $mid = 0;

        while ($lo < $hi) {
            $mid = $lo + (($hi - $lo) >> 1);
            switch (true) {
            case $key < $arr[$mid]:
                $hi = $mid;
                break;
            case $key > $arr[$mid]:
                $lo = $mid + 1;
                break;
            default:
                return $mid;
            }
        }

        return count($arr);
    }
?>
