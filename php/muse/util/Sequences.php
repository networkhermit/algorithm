<?php
    namespace muse\util\Sequences;

    require_once "muse/algorithms/sorting/QuickSort.php";
    require_once "muse/util/RandomFactory.php";

    use muse\algorithms\sorting\QuickSort;
    use muse\util\RandomFactory;

    function inspect(array &$arr): void {
        print("[\n");
        foreach ($arr as $i => $v) {
            printf("\t#%04X  -->  %d\n", $i, $v);
        }
        print("]\n");
    }

    function isSorted(array &$arr): bool {
        for ($i = 1, $length = count($arr); $i < $length; $i++) {
            if ($arr[$i] < $arr[$i - 1]) {
                return false;
            }
        }

        return true;
    }

    function parityChecksum(array &$arr): int {
        $checksum = 0;

        foreach ($arr as $v) {
            $checksum ^= $v;
        }

        return $checksum;
    }

    function reverse(array &$arr): void {
        $k;
        $temp;

        for ($i = 0, $bound = count($arr) >> 1, $length = count($arr); $i < $bound; $i++) {
            $k = $length - $i - 1;
            $temp = $arr[$i];
            $arr[$i] = $arr[$k];
            $arr[$k] = $temp;
        }
    }

    function shuffle(array &$arr): void {
        $k;
        $temp;

        RandomFactory\seed();
        for ($i = 0, $length = count($arr); $i < $length; $i++) {
            $k = RandomFactory\genIntN($i, $length);
            $temp = $arr[$i];
            $arr[$i] = $k;
            $k = $temp;
        }
    }

    function sort(array &$arr): void {
        QuickSort\sort($arr);
    }
?>
