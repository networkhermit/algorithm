<?php
    require_once "vac.fun/algorithms/search/BinarySearch.php";
    require_once "vac.fun/util/SequenceBuilder.php";
    require_once "vac.fun/util/TestRunner.php";

    use vac\fun\algorithms\search\BinarySearch;
    use vac\fun\util\SequenceBuilder;
    use vac\fun\util\TestRunner;

    function testBinarySearch(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packIncreasing($arr);

        if (BinarySearch\find($arr, -1) != $size) {
            return false;
        }

        if (BinarySearch\find($arr, 2147483647) != $size) {
            return false;
        }

        foreach ($arr as $i => $v) {
            if (BinarySearch\find($arr, $v) != $i) {
                return false;
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testBinarySearch());
    }
?>
