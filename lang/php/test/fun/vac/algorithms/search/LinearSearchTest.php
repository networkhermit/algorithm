<?php
    require_once "fun/vac/algorithms/search/LinearSearch.php";
    require_once "fun/vac/util/SequenceBuilder.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\algorithms\search\LinearSearch;
    use fun\vac\util\SequenceBuilder;
    use fun\vac\util\TestRunner;

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
