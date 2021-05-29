<?php
    namespace vac\fun\algorithms\sorting\MergeSort;

    function merge(array &$arr, int $lo, int $mid, int $hi): void {
        if ($lo == $mid) {
            return;
        }

        merge($arr, $lo, ($lo + $mid) >> 1, $mid);
        merge($arr, $mid, ($mid + $hi) >> 1, $hi);

        $m = $lo;
        $n = $mid;

        $sorted = array_pad(array(), $hi - $lo, 0);

        foreach ($sorted as &$v) {
            if ($m != $mid && ($n == $hi || $arr[$m] < $arr[$n])) {
                $v = $arr[$m];
                $m++;
            } else {
                $v = $arr[$n];
                $n++;
            }
        }

        $cursor = 0;
        for ($i = $lo; $i < $hi; $i++) {
            $arr[$i] = $sorted[$cursor];
            $cursor++;
        }
    }

    function sort(array &$arr): void {
        merge($arr, 0, count($arr) >> 1, count($arr));
    }
?>
