<?php
    require_once "fun/vac/algorithms/search/BinarySearch.php";
    require_once "fun/vac/util/SequenceBuilder.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\algorithms\search\BinarySearch;
    use fun\vac\util\SequenceBuilder;
    use fun\vac\util\TestRunner;

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
