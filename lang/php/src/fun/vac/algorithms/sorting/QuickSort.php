<?php
    namespace fun\vac\algorithms\sorting\QuickSort;

    function partition(array &$arr, int $lo, int $hi): void {
        if ($lo === $hi) {
            return;
        }

        $pivot = $arr[$lo];

        $left  = $lo;
        $right = $hi - 1;

        while ($left != $right) {
            for ($i = $right; $i > $left; $i--) {
                if ($arr[$i] < $pivot) {
                    $arr[$left] = $arr[$i];
                    $arr[$i] = $pivot;
                    break;
                }
                $right -= 1;
            }
            for ($i = $left; $i < $right; $i++) {
                if ($arr[$i] > $pivot) {
                    $arr[$right] = $arr[$i];
                    $arr[$i] = $pivot;
                    break;
                }
                $left += 1;
            }
        }

        partition($arr, $lo, $left);
        partition($arr, $left + 1, $hi);
    }

    function sort(array &$arr): void {
        partition($arr, 0, count($arr));
    }
?>
