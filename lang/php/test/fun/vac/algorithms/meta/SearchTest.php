<?php
    require_once "fun/vac/algorithms/meta/Search.php";
    require_once "fun/vac/util/SequenceBuilder.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\algorithms\meta\Search;
    use fun\vac\util\SequenceBuilder;
    use fun\vac\util\TestRunner;

    function testBinarySearch(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packIncreasing($arr);

        if (Search\binarySearch($arr, -1) != $size) {
            return false;
        }

        if (Search\binarySearch($arr, 2147483647) != $size) {
            return false;
        }

        foreach ($arr as $i => $v) {
            if (Search\binarySearch($arr, $v) != $i) {
                return false;
            }
        }

        return true;
    }

    function testLinearSearch(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packIncreasing($arr);

        if (Search\linearSearch($arr, -1) != $size) {
            return false;
        }

        if (Search\linearSearch($arr, 2147483647) != $size) {
            return false;
        }

        foreach ($arr as $i => $v) {
            if (Search\linearSearch($arr, $v) != $i) {
                return false;
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {

        TestRunner\parseTest(testBinarySearch());

        TestRunner\parseTest(testLinearSearch());
    }
?>
