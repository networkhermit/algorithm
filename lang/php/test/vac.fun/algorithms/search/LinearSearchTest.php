<?php
    require_once "vac.fun/algorithms/search/LinearSearch.php";
    require_once "vac.fun/util/SequenceBuilder.php";
    require_once "vac.fun/util/TestRunner.php";

    use vac\fun\algorithms\search\LinearSearch;
    use vac\fun\util\SequenceBuilder;
    use vac\fun\util\TestRunner;

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
