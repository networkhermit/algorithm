<?php
    namespace fun\vac\algorithms\search\BinarySearch;

    function find(array &$arr, $key): int {
        $lo = 0;
        $hi = count($arr);

        $mid = 0;

        while ($lo < $hi) {
            $mid = $lo + (($hi - $lo) >> 1);
            if ($key < $arr[$mid]) {
                $hi = $mid;
            } else if ($key > $arr[$mid]) {
                $lo = $mid + 1;
            } else {
                return $mid;
            }
        }

        return count($arr);
    }
?>
