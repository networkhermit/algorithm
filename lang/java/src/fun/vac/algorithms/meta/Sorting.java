package fun.vac.algorithms.meta;

import fun.vac.algorithms.sorting.BubbleSort;
import fun.vac.algorithms.sorting.InsertionSort;
import fun.vac.algorithms.sorting.MergeSort;
import fun.vac.algorithms.sorting.QuickSort;
import fun.vac.algorithms.sorting.SelectionSort;

public class Sorting {

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
