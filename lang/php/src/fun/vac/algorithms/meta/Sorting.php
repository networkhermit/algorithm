<?php
    namespace fun\vac\algorithms\meta\Sorting;

    require_once "fun/vac/algorithms/sorting/BubbleSort.php";
    require_once "fun/vac/algorithms/sorting/InsertionSort.php";
    require_once "fun/vac/algorithms/sorting/MergeSort.php";
    require_once "fun/vac/algorithms/sorting/QuickSort.php";
    require_once "fun/vac/algorithms/sorting/SelectionSort.php";

    use fun\vac\algorithms\sorting\BubbleSort;
    use fun\vac\algorithms\sorting\InsertionSort;
    use fun\vac\algorithms\sorting\MergeSort;
    use fun\vac\algorithms\sorting\QuickSort;
    use fun\vac\algorithms\sorting\SelectionSort;

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
