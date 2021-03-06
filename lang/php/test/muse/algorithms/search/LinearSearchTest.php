<?php
    require_once "muse/algorithms/search/LinearSearch.php";
    require_once "muse/util/SequenceBuilder.php";
    require_once "muse/util/TestRunner.php";

    use muse\algorithms\search\LinearSearch;
    use muse\util\SequenceBuilder;
    use muse\util\TestRunner;

    function testLinearSearch(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packIncreasing($arr);

        if (LinearSearch\find($arr, -1) != $size) {
            return false;
        }

        if (LinearSearch\find($arr, 2147483647) != $size) {
            return false;
        }

        foreach ($arr as $i => $v) {
            if (LinearSearch\find($arr, $v) != $i) {
                return false;
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testLinearSearch());
    }
?>
