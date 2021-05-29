<?php
    require_once "vac.fun/algorithms/meta/Sorting.php";
    require_once "vac.fun/util/SequenceBuilder.php";
    require_once "vac.fun/util/Sequences.php";
    require_once "vac.fun/util/TestRunner.php";

    use vac\fun\algorithms\meta\Sorting;
    use vac\fun\util\SequenceBuilder;
    use vac\fun\util\Sequences;
    use vac\fun\util\TestRunner;

    function testBubbleSort(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packRandom($arr);

        $checksum = Sequences\parityChecksum($arr);

        Sorting\bubbleSort($arr);

        if (Sequences\parityChecksum($arr) != $checksum) {
            return false;
        }

        return Sequences\isSorted($arr);
    }

    function testInsertionSort(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packRandom($arr);

        $checksum = Sequences\parityChecksum($arr);

        Sorting\insertionSort($arr);

        if (Sequences\parityChecksum($arr) != $checksum) {
            return false;
        }

        return Sequences\isSorted($arr);
    }

    function testMergeSort(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packRandom($arr);

        $checksum = Sequences\parityChecksum($arr);

        Sorting\mergeSort($arr);

        if (Sequences\parityChecksum($arr) != $checksum) {
            return false;
        }

        return Sequences\isSorted($arr);
    }

    function testQuickSort(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packRandom($arr);

        $checksum = Sequences\parityChecksum($arr);

        Sorting\quickSort($arr);

        if (Sequences\parityChecksum($arr) != $checksum) {
            return false;
        }

        return Sequences\isSorted($arr);
    }

    function testSelectionSort(): bool {
        $size = 32768;

        $arr = array_pad(array(), $size, 0);
        SequenceBuilder\packRandom($arr);

        $checksum = Sequences\parityChecksum($arr);

        Sorting\selectionSort($arr);

        if (Sequences\parityChecksum($arr) != $checksum) {
            return false;
        }

        return Sequences\isSorted($arr);
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testBubbleSort());

        TestRunner\parseTest(testInsertionSort());

        TestRunner\parseTest(testMergeSort());

        TestRunner\parseTest(testQuickSort());

        TestRunner\parseTest(testSelectionSort());
    }
?>