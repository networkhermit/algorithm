<?php
    declare(strict_types=1);

    namespace muse\algorithms\meta\Sorting;

    require_once "muse/algorithms/sorting/BubbleSort.php";
    require_once "muse/algorithms/sorting/InsertionSort.php";
    require_once "muse/algorithms/sorting/MergeSort.php";
    require_once "muse/algorithms/sorting/QuickSort.php";
    require_once "muse/algorithms/sorting/SelectionSort.php";

    use muse\algorithms\sorting\BubbleSort;
    use muse\algorithms\sorting\InsertionSort;
    use muse\algorithms\sorting\MergeSort;
    use muse\algorithms\sorting\QuickSort;
    use muse\algorithms\sorting\SelectionSort;

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
