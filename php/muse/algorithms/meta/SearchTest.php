<?php
    require_once "muse/algorithms/meta/Search.php";
    require_once "muse/util/SequenceBuilder.php";
    require_once "muse/util/TestRunner.php";

    use muse\algorithms\meta\Search;
    use muse\util\SequenceBuilder;
    use muse\util\TestRunner;

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
        TestRunner\pick(testBinarySearch);

        TestRunner\pick(testLinearSearch);
    }
?>
