<?php
    require_once "vac.fun/algorithms/sorting/SelectionSort.php";
    require_once "vac.fun/util/SequenceBuilder.php";
    require_once "vac.fun/util/Sequences.php";
    require_once "vac.fun/util/TestRunner.php";

    use vac\fun\algorithms\sorting\SelectionSort;
    use vac\fun\util\SequenceBuilder;
    use vac\fun\util\Sequences;
    use vac\fun\util\TestRunner;

    function testSelectionSort(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packRandom($arr);

        $checksum = Sequences\parityChecksum($arr);

        SelectionSort\sort($arr);

        if (Sequences\parityChecksum($arr) != $checksum) {
            return false;
        }

        return Sequences\isSorted($arr);
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testSelectionSort());
    }
?>
