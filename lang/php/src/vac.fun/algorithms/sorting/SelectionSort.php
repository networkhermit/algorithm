<?php
    namespace vac\fun\algorithms\sorting\SelectionSort;

    function sort(array &$arr): void {
        if (count($arr) == 0) {
            return;
        }

        $temp = NULL;

        $iMin = 0;

        for ($i = 0, $bound = count($arr) - 1; $i < $bound; $i++) {
            $iMin = $i;
            for ($j = $i + 1; $j <= $bound; $j++) {
                if ($arr[$j] < $arr[$iMin]) {
                    $iMin = $j;
                }
            }
            if ($iMin != $i) {
                $temp = $arr[$i];
                $arr[$i] = $arr[$iMin];
                $arr[$iMin] = $temp;
            }
        }
    }
?>
