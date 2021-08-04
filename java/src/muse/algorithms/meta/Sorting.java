package muse.algorithms.meta;

import muse.algorithms.sorting.BubbleSort;
import muse.algorithms.sorting.InsertionSort;
import muse.algorithms.sorting.MergeSort;
import muse.algorithms.sorting.QuickSort;
import muse.algorithms.sorting.SelectionSort;

public final class Sorting {

    private Sorting() {}

    public static <T extends Comparable<T>> void bubbleSort(T[] arr) {
        BubbleSort.sort(arr);
    }

    public static <T extends Comparable<T>> void insertionSort(T[] arr) {
        InsertionSort.sort(arr);
    }

    public static <T extends Comparable<T>> void mergeSort(T[] arr) {
        MergeSort.sort(arr);
    }

    public static <T extends Comparable<T>> void quickSort(T[] arr) {
        QuickSort.sort(arr);
    }

    public static <T extends Comparable<T>> void selectionSort(T[] arr) {
        SelectionSort.sort(arr);
    }
}
