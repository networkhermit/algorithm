<?php
    namespace fun\vac\algorithms\sorting\BubbleSort;

    function sort(array &$arr): void {
        $temp = NULL;

        $margin = 0;
        $unsorted = count($arr);

        while ($unsorted > 1) {
            $margin = 0;
            for ($i = 1; $i < $unsorted; $i++) {
                if ($arr[$i - 1] > $arr[$i]) {
                    $temp = $arr[$i - 1];
                    $arr[$i - 1] = $arr[$i];
                    $arr[$i] = $temp;
                    $margin = $i;
                }
            }
            $unsorted = $margin;
        }
    }
?>
