require "vac.fun/algorithms/sorting/BubbleSort"
require "vac.fun/algorithms/sorting/InsertionSort"
require "vac.fun/algorithms/sorting/MergeSort"
require "vac.fun/algorithms/sorting/QuickSort"
require "vac.fun/algorithms/sorting/SelectionSort"

module Sorting

    def self.bubbleSort(arr)
        BubbleSort.sort(arr)
    end

    def self.insertionSort(arr)
        InsertionSort.sort(arr)
    end

    def self.mergeSort(arr)
        MergeSort.sort(arr)
    end

    def self.quickSort(arr)
        QuickSort.sort(arr)
    end

    def self.selectionSort(arr)
        SelectionSort.sort(arr)
    end
end
