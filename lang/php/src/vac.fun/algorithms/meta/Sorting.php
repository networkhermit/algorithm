<?php
    namespace vac\fun\algorithms\meta\Sorting;

    require_once "vac.fun/algorithms/sorting/BubbleSort.php";
    require_once "vac.fun/algorithms/sorting/InsertionSort.php";
    require_once "vac.fun/algorithms/sorting/MergeSort.php";
    require_once "vac.fun/algorithms/sorting/QuickSort.php";
    require_once "vac.fun/algorithms/sorting/SelectionSort.php";

    use vac\fun\algorithms\sorting\BubbleSort;
    use vac\fun\algorithms\sorting\InsertionSort;
    use vac\fun\algorithms\sorting\MergeSort;
    use vac\fun\algorithms\sorting\QuickSort;
    use vac\fun\algorithms\sorting\SelectionSort;

    function bubbleSort(array &$arr): void {
        BubbleSort\sort($arr);
    }

    function insertionSort(array &$arr): void {
        InsertionSort\sort($arr);
    }

    function mergeSort(array &$arr): void {
        MergeSort\sort($arr);
    }

    function quickSort(array &$arr): void {
        QuickSort\sort($arr);
    }

    function selectionSort(array &$arr): void {
        SelectionSort\sort($arr);
    }
?>
