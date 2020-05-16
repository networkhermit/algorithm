<?php
    require_once "fun/vac/algorithms/sorting/QuickSort.php";
    require_once "fun/vac/util/SequenceBuilder.php";
    require_once "fun/vac/util/Sequences.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\algorithms\sorting\QuickSort;
    use fun\vac\util\SequenceBuilder;
    use fun\vac\util\Sequences;
    use fun\vac\util\TestRunner;

    function testQuickSort(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packRandom($arr);

        $checksum = Sequences\parityChecksum($arr);

        QuickSort\sort($arr);

        if (Sequences\parityChecksum($arr) != $checksum) {
            return false;
        }

        return Sequences\isSorted($arr);
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testQuickSort());
    }
?>
